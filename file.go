package gotool

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	PermFile fs.FileMode = 0666
	PermDir              = 0755
)

// file is exists
// 判断地址是否存在
func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// read file content to map instance
func FileReadToMap(path string) (map[string]interface{}, error) {
	var fileData map[string]interface{}
	err := FileReadForInterface(path, &fileData)
	if err != nil {
		return nil, err
	}
	return fileData, nil
}

func FileReadForInterface(path string, i interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, i)
}

func FileWriteWithInterface(path string, data interface{}) error {
	var writeBytes []byte
	var err error
	switch data.(type) {
	case string:
		writeBytes = []byte(data.(string))
	case []byte:
		writeBytes = data.([]byte)
	default:
		writeBytes, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(path, writeBytes, PermFile)
}

// 列举文件列表
// root: 遍历的目录
// hasHide: 是否包含隐藏文件
// isRecursion: 是否递归子文件夹
// fn: 需要执行的方法
func FileList(root string, hasHide bool, isRecursion bool, fn filepath.WalkFunc) error {
	var err error
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// 不处理文件夹
			if info.IsDir() {
				return nil
			}
			dirName := filepath.Base(filepath.Dir(path))
			// 判断是否处理隐藏文件
			if (strings.HasPrefix(info.Name(), ".") || strings.HasPrefix(dirName, ".")) && !hasHide {
				return nil
			}
			// 判断是否递归处理
			if path != filepath.Join(root, info.Name()) && !isRecursion {
				return nil
			}
			return fn(path, info, err)
		})
	return err
}

func FilesRemove(paths []string) error {
	errs := make([]string, 0)
	for _, p := range paths {
		err := os.Remove(p)
		if err != nil {
			errs = append(errs, fmt.Sprintf("%s: %v", p, err))
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	} else {
		return nil
	}
}

// 自动对下载地址进行重命名
// examples:
// ~/download/main.go
// ~/download/main(1).go
// ~/download/main(2).go
func FileAutoReDownloadName(path string) string {
	if !FileExists(path) {
		return path
	}
	i := 1
	var newPath string
	for true {
		ext := strings.TrimLeft(filepath.Ext(path), ".")
		if ext != "" {

			prefix := strings.TrimRight(strings.TrimRight(path, ext), ".")
			newPath = fmt.Sprintf("%s(%d).%s", prefix, i, ext)
		} else {
			newPath = fmt.Sprintf("%s(%d)", path, i)
		}
		if !FileExists(newPath) {
			return newPath
		}
		i++
	}
	return path
}
