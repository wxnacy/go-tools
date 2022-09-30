package gotool

import "testing"

func TestFormatSize(t *testing.T) {
	var res string
	res = FormatSize(123)
	if res != "123B" {
		t.Error(res)
	}
	res = FormatSize(1024)
	if res != "1K" {
		t.Error(res)
	}
	res = FormatSize(1124)
	if res != "1.10K" {
		t.Error(res)
	}
	res = FormatSize(1124123)
	if res != "1.07M" {
		t.Error(res)
	}
	res = FormatSize(999999999999999999)
	if res != "999999999999999999B" {
		t.Error(res)
	}
}

func TestFormatFloat(t *testing.T) {
	var res float64
	res = FormatFloat(1.234, 2)
	if res != 1.23 {
		t.Error(res)
	}
}
