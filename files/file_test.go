package files

import "testing"

func TestFileExists(t *testing.T) {
	if !FileExists("file.go") {
		t.Error("file.go")
	}
}
