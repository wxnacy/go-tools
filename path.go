package tools

import (
	"os"
	"path/filepath"
)

// 地址所在的目录是否存在
func PathDirExists(path string) bool {
	dir := filepath.Dir(path)
	return DirExists(dir)
}

// 地址是否存在，不区分文件和目录
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
