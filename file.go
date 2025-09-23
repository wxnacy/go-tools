package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	PermFile fs.FileMode = 0666
	PermDir              = 0755

	FileSegmentSize = 8 * 1024 * 1024
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
		if !PathExists(newPath) {
			return newPath
		}
		i++
	}
	return path
}

// 合并文件
func FilesMerge(target string, sources []string, perm fs.FileMode) error {
	writeFile, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	defer writeFile.Close()
	if err != nil {
		return err
	}

	appendFile := func(path string) error {
		tempFile, err := os.Open(path)
		defer tempFile.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(writeFile, tempFile)
		return err
	}

	for _, path := range sources {
		err = appendFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}

// 复制文件
// src string: 源文件
// dst string: 目标文件
// segmentSize int64: 分片大小
func FileCopy(src, dst string, segmentSize int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		return err
	}

	buf := make([]byte, segmentSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}
