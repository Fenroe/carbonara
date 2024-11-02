package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func readPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	// Switch terminal into raw mode
	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(syscall.Stdin), oldState)

	var bytePassword []byte
	for {
		// Read a single byte from the input
		var buf [1]byte
		n, err := os.Stdin.Read(buf[:])
		if err != nil || n == 0 {
			break
		}

		// Check for newline or carriage return (Enter key)
		if buf[0] == '\n' || buf[0] == '\r' {
			fmt.Println() // move to a new line
			break
		}

		// Handle backspace (delete the last character if exists)
		if buf[0] == 127 || buf[0] == '\b' {
			if len(bytePassword) > 0 {
				bytePassword = bytePassword[:len(bytePassword)-1]
				fmt.Print("\b \b") // erase the last asterisk
			}
		} else {
			// Add the character to the password slice
			bytePassword = append(bytePassword, buf[0])
			// Print asterisk as a mask
			fmt.Print("*")
		}
	}
	password := string(bytePassword)
	err = validatePassword(password)
	if err != nil {
		return "", err
	}
	return string(bytePassword), err
}
