package main

import (
	"FileTools/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.FileToBase64Cmd.Execute(); err != nil {
		fmt.Printf("file to base64 fail: %v", err)
		os.Exit(-1)
	}
}
