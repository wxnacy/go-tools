package gotool

import (
	"fmt"
	"os"
	"path/filepath"
)

// dir path is exists
func DirExists(dirpath string) bool {
	info, err := os.Stat(dirpath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// 判断目录是否存在，否则创建
func DirExistsOrCreate(dirpath string) error {
	info, err := os.Stat(dirpath)
	if os.IsNotExist(err) {
		return os.MkdirAll(dirpath, PermDir)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is exists but is not dir", dirpath)
	}
	return nil
}

// Get dir file total size
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// Get dir file total size and format
func DirSizeFormat(path string) (string, error) {
	size, err := DirSize(path)
	if err != nil {
		return "", err
	}
	return FormatSize(size), nil
}
