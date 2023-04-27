package gotool

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
