package tools

import "testing"

func TestDirExists(t *testing.T) {
	if !DirExists("/tmp") {
		t.Error("/tmp is dir")
	}
	if DirExists("file.go") {
		t.Error("file.go is not dir")
	}
}
