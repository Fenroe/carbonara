package main

import (
	"fmt"

	"golang.design/x/clipboard"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Reading your clipboard...")
	err := clipboard.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	text := string(clipboard.Read(clipboard.FmtText))
	if text != "" {
		fmt.Println(text)
		return
	}
}
