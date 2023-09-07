package tools

import (
	"strings"
	"testing"
)

func TestStringFromReader(t *testing.T) {
	str := "wxnacy"
	r := strings.NewReader(str)
	s, err := StringFromReader(r)
	if s != str {
		t.Errorf("%s ! %s", str, s)
	}
	if err != nil {
		t.Errorf("err is %v", err)
	}
}

func TestStringBackspace(t *testing.T) {
	var in, out string
	for _, line := range [][]string{
		[]string{"wxnacy", "wxnac"},
		[]string{"你好", "你"},
		[]string{"wxnacy 你好", "wxnacy 你"},
		[]string{"", ""},
	} {
		in = line[0]
		out = line[1]
		res := StringBackspace(in)
		if res != out {
			t.Errorf("StringBackspace %s == %s error %s", in, out, res)
		}
	}
}
