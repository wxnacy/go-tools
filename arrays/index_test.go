package arrays

import "testing"

func TestIndex(t *testing.T) {
	var arr = []string{"wxnacy", "winn"}
	var s = "wxnacy"
	i := Index(arr, s)
	if i != 0 {
		t.Error(i)
	}
}

func TestStringIndex(t *testing.T) {
	var arr = []string{"wxnacy", "winn"}
	var s = "wxnacy"
	i := StringIndex(arr, s)
	if i != 0 {
		t.Error(i)
	}
}

func TestIntIndex(t *testing.T) {
	var arr = []int{1, 3, 4, 8, 12, 4, 9}
	var s = 12
	i := IntIndex(arr, s)
	if i != 4 {
		t.Error(i)
	}
	i = IntIndex(arr, s)
	if i != 4 {
		t.Error(i)
	}
}
