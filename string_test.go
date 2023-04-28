package gotool

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
