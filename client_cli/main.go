package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"golang.design/x/clipboard"
)

func main() {
	godotenv.Load()
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
