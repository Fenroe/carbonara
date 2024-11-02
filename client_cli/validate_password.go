package main

import (
	"errors"
	"fmt"
)

func validatePassword(password string) error {
	minPasswordLength := 7
	if len(password) < minPasswordLength {
		errorMsg := fmt.Sprintf("password must contain at least %d characters", minPasswordLength)
		return errors.New(errorMsg)
	}
	return nil
}
