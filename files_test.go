package gotool

import "testing"

func TestFileExists(t *testing.T) {
	if !FileExists("files.go") {
		t.Error("files.go not found")
	}
}

func TestDirExists(t *testing.T) {
	if !DirExists("main") {
		t.Error("main")
	}
}
