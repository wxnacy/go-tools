package arrays

import "testing"

func TestContains(t *testing.T) {
	var arr = []string{"wxnacy", "winn"}
	var s = "wxnacy"
	i := Contains(arr, s)
	if !i {
		t.Error(i)
	}
}

func TestStringContains(t *testing.T) {
	var arr = []string{"wxnacy", "winn"}
	var s = "wxnacy"
	i := StringContains(arr, s)
	if !i {
		t.Error(i)
	}
}

func TestIntContains(t *testing.T) {
	var arr = []int{1, 3, 4, 8, 12, 4, 9}
	var s = 12
	i := IntContains(arr, s)
	if !i {
		t.Error(i)
	}
	i = IntContains(arr, s)
	if !i {
		t.Error(i)
	}
}
