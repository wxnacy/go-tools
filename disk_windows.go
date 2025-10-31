//go:build windows
// +build windows

package tools

import (
	"golang.org/x/sys/windows"
)

// 磁盘剩余空间
func DiskFree(path string) (uint64, error) {
	// 将路径转换为 Windows 系统调用所需的 UTF16 编码
	pathPtr, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return 0, err
	}

	var freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes uint64
	// 调用 Windows API 获取磁盘空间信息
	err = windows.GetDiskFreeSpaceEx(pathPtr, &freeBytesAvailable, &totalNumberOfBytes, &totalNumberOfFreeBytes)
	if err != nil {
		return 0, err
	}

	// 返回可用空间字节数
	return freeBytesAvailable, nil
}