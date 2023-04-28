package tools

import "testing"

func TestIndex(t *testing.T) {
	var arr = []string{"wxnacy", "winn"}
	var s = "wxnacy"
	i := ArrayIndex(arr, s)
	if i != 0 {
		t.Error(i)
	}
}

func TestStringIndex(t *testing.T) {
	var arr = []string{"wxnacy", "winn"}
	var s = "wxnacy"
	i := ArrayIndexString(arr, s)
	if i != 0 {
		t.Error(i)
	}
}

func TestIntIndex(t *testing.T) {
	var arr = []int{1, 3, 4, 8, 12, 4, 9}
	var s = 12
	i := ArrayIndexInt(arr, s)
	if i != 4 {
		t.Error(i)
	}
}

func TestIndexBool(t *testing.T) {
	var arr = []bool{true, false, false}
	var s = true
	i := ArrayIndexBool(arr, s)
	if i != 0 {
		t.Error(i)
	}
}

func TestContains(t *testing.T) {
	var arr = []string{"wxnacy", "winn"}
	var s = "wxnacy"
	i := ArrayContains(arr, s)
	if !i {
		t.Error(i)
	}
}

func TestStringContains(t *testing.T) {
	var arr = []string{"wxnacy", "winn"}
	var s = "wxnacy"
	i := ArrayContainsString(arr, s)
	if !i {
		t.Error(i)
	}
}

func TestIntContains(t *testing.T) {
	var arr = []int{1, 3, 4, 8, 12, 4, 9}
	var s = 12
	i := ArrayContainsInt(arr, s)
	if !i {
		t.Error(i)
	}
}

func TestContainsBool(t *testing.T) {
	var arr = []bool{false, true, false}
	var s = true
	i := ArrayContainsBool(arr, s)
	if !i {
		t.Error(i)
	}
}
