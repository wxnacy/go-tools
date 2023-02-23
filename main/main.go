package main

import (
	"fmt"
	"gotool"
	"gotool/files"
)

func main() {
	fmt.Println(gotool.FormatSize(123))
	fmt.Println(gotool.FormatSize(1024))
	fmt.Println(gotool.FormatSize(1124))
	fmt.Println(gotool.FormatSize(1124123))
	fmt.Println(gotool.FormatSize(11241234300))
	fmt.Println(gotool.FormatSize(9091241234300))
	a := 12.0
	// b := 12.3
	fmt.Println(float64(int64(a)) - a)
	fmt.Println(files.DirSizeFormat("/Users/wxnacy/Downloads/test"))
}
