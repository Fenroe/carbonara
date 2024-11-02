package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func readEmail(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	email, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("couldn't get email")
	}
	email = strings.ReplaceAll(email, "\n", "")
	err = validateEmail(email)
	if err != nil {
		return "", err
	}
	return email, err
}
