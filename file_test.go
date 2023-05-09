package tools

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestFileExists(t *testing.T) {
	if !FileExists("file.go") {
		t.Error("file.go not found")
	}
}

func TestFileAutoReDownloadName(t *testing.T) {
	os.Chdir(cacheDir)
	var path, newPath, name string
	path = IDGen()
	path = FileAutoReDownloadName(path)
	if path != path {
		t.Errorf("%s == %s\n", path, path)
	}
	FileWriteWithInterface(path, "w")
	newPath = FileAutoReDownloadName(path)
	if newPath != path+"(1)" {
		t.Errorf("%s == %s\n", path+"(1)", newPath)
	}
	// xxx.mp4
	name = IDGen()
	path = name + ".mp4"
	path = FileAutoReDownloadName(path)
	if path != path {
		t.Errorf("%s == %s\n", path, path)
	}
	FileWriteWithInterface(path, "w")
	newPath = FileAutoReDownloadName(path)
	if newPath != name+"(1)"+".mp4" {
		t.Errorf("%s == %s\n", name+"(1)"+".mp4", newPath)
	}
	// xxx.mp4.mp4
	name = IDGen() + ".mp4"
	path = name + ".mp4"
	path = FileAutoReDownloadName(path)
	if path != path {
		t.Errorf("%s == %s\n", path, path)
	}
	FileWriteWithInterface(path, "w")
	newPath = FileAutoReDownloadName(path)
	if newPath != name+"(1)"+".mp4" {
		t.Errorf("%s == %s\n", name+"(1)"+".mp4", newPath)
	}

	DirFilesRemove(cacheDir, "")
}

func TestFilesRemove(t *testing.T) {
	dir, _ := os.Getwd()
	defer os.Chdir(dir)
	os.Chdir(cacheDir)
	paths := make([]string, 0)
	for i := 0; i < 5; i++ {
		path := "remove_" + strconv.Itoa(rand.Intn(1000))
		FileWriteWithInterface(path, "w")
		paths = append(paths, path)
	}
	FilesRemove(paths)
	for _, p := range paths {
		if FileExists(p) {
			t.Errorf("%s exist", p)
		}
	}
}

func TestFilesMerge(t *testing.T) {
	os.Chdir(cacheDir)
	paths := make([]string, 0)
	for i := 0; i < 5; i++ {
		path := "merge_" + strconv.Itoa(rand.Intn(1000))
		FileWriteWithInterface(path, "w")
		paths = append(paths, path)
	}
	FilesMerge("merge", paths, PermFile)
	b, _ := ioutil.ReadFile("merge")
	if string(b) != "wwwww" {
		t.Errorf("%s != wwwww", string(b))
	}
	DirFilesRemove(cacheDir, "")
}
