package tools

import (
	"os"
	"testing"
)

func TestMd5(t *testing.T) {
	var res string
	res = Md5("wxnacy")
	if res != "1f806eb48b670c40af49a3f764ba086f" {
		t.Errorf("%s is error", res)
	}
}

func TestMd5File(t *testing.T) {
	os.Chdir(cacheDir)
	var res string
	FileWriteWithInterface("TestMd5File", "wxnacy")
	res, err := Md5File("TestMd5File")
	if err != nil {
		t.Error(err)
	}
	if res != "1f806eb48b670c40af49a3f764ba086f" {
		t.Errorf("%s != 1f806eb48b670c40af49a3f764ba086f", res)
	}
	DirFilesRemove(cacheDir, "")
}
