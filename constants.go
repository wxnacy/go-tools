package tools

import "path/filepath"

const (
	cacheDir = "/tmp/go-tools"
)

func init() {
	DirExistsOrCreate(cacheDir)
}

func joinCachePath(elem ...string) string {
	return join(cacheDir, elem...)
}

func join(root string, elem ...string) string {
	elem = append([]string{root}, elem...)
	return filepath.Join(elem...)
}
