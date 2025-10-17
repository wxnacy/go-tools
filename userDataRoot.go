package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// 获取用户数据目录（遵循XDG规范）
func GetUserDataRoot() (string, error) {
	// 1. 优先检查XDG_DATA_HOME环境变量（Linux/macOS）
	if xdgDataHome := os.Getenv("XDG_DATA_HOME"); xdgDataHome != "" {
		return xdgDataHome, nil
	}

	// 2. 未设置则使用系统默认路径
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("无法获取用户主目录: %w", err)
	}

	// 根据操作系统拼接默认路径
	switch os := runtime.GOOS; os {
	case "linux":
		return filepath.Join(homeDir, ".local", "share"), nil
	case "darwin": // macOS
		return filepath.Join(homeDir, "Library", "Application Support"), nil
	case "windows":
		return filepath.Join(homeDir, "AppData", "Local"), nil
	default:
		return "", fmt.Errorf("不支持的操作系统: %s", os)
	}
}
