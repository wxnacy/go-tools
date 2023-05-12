package tools

import "fmt"

func FuncConfirm(msg string) bool {
	var input string
	fmt.Printf("%s(y/N): ", msg)
	fmt.Scanln(&input)
	if input == "y" {
		return true
	} else {
		return false
	}
}
