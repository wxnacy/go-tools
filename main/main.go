package main

import (
	"fmt"

	"github.com/wxnacy/gotool"
	"github.com/wxnacy/gotool/files"
)

func main() {
	fmt.Println(gotool.FormatSize(123))
	fmt.Println(gotool.FormatSize(1024))
	fmt.Println(gotool.FormatSize(1124))
	fmt.Println(gotool.FormatSize(1124123))
	fmt.Println(gotool.FormatSize(11241234300))
	fmt.Println(gotool.FormatSize(9091241234300))
	a := 12.1
	fmt.Println(int64(a) == 12.0)
	fmt.Println(files.DirSizeFormat("/Users/wxnacy/Downloads/test"))
}
