package tools

import (
	"io"
	"strings"
)

// 通过 io.Reader 转为 string
func StringFromReader(r io.Reader) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// 字符串进行退格
func StringBackspace(s string) string {
	if s == "" {
		return s
	}
	total := 0
	for range s {
		total++
	}
	res := ""
	backIndex := 0
	for _, r := range s {
		backIndex++
		if backIndex < total {
			res += string(r)
		}
	}
	return res
}
