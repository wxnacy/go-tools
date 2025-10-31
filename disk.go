//go:build !windows
// +build !windows

package tools

import "golang.org/x/sys/unix"

// 磁盘剩余空间
func DiskFree(path string) (uint64, error) {
	fs := unix.Statfs_t{}
	err := unix.Statfs(path, &fs)
	if err != nil {
		return 0, err
	}
	return fs.Bfree * uint64(fs.Bsize), nil
}
