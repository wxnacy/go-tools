package main

import (
	"fmt"
	"os"

	tools "github.com/wxnacy/go-tools"
)

// go run examples/download/main.go https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png tmp/google.png
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <url> <filepath>")
		return
	}
	url := os.Args[1]
	filepath := os.Args[2]
	err := tools.Download(url, filepath)
	if err != nil {
		fmt.Printf("Error downloading file: %s\n", err)
		return
	}
	fmt.Printf("File downloaded successfully to %s\n", filepath)
}
