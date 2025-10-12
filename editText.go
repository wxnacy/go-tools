package tools

import (
	"fmt"
	"os"
	"os/exec"
)

// 在编辑器中编辑文本
func EditTextInEditer(editer, text string) (string, error) {
	// 创建一个临时文件
	tmpfile, err := os.CreateTemp("", "edit-*.txt")
	if err != nil {
		return "", fmt.Errorf("could not create temp file: %w", err)
	}
	// 使用 defer 来确保临时文件在函数结束时被删除
	defer os.Remove(tmpfile.Name())

	// 将初始文本写入临时文件
	if _, err := tmpfile.Write([]byte(text)); err != nil {
		return "", fmt.Errorf("could not write to temp file: %w", err)
	}
	if err := tmpfile.Close(); err != nil {
		return "", fmt.Errorf("could not close temp file: %w", err)
	}

	// 构建运行 nvim 的命令
	cmd := exec.Command(editer, tmpfile.Name())

	// 将子进程的 stdin, stdout, stderr 连接到当前进程
	// 这样用户就可以直接在终端中与 nvim 交互
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 运行命令并等待它完成
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running nvim: %w", err)
	}

	// 读取编辑后文件的内容
	editedContent, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		return "", fmt.Errorf("could not read temp file after editing: %w", err)
	}

	return string(editedContent), nil
}
