package gotool

import "testing"

func TestMd5(t *testing.T) {
	var res string
	res = Md5("wxnacy")
	if res != "1f806eb48b670c40af49a3f764ba086f" {
		t.Errorf("%s is error", res)
	}
}

func TestMd5File(t *testing.T) {
	var res string
	res, err := Md5File("./LICENSE")
	if err != nil {
		t.Error(err)
	}
	if res != "ecd2834adf1e5b6e960e4f82351d309d" {
		t.Errorf("%s is error", res)
	}
}
