package idiomatic

import (
	"errors"
	"fmt"
)

func DoSomething(input int) (string, error) {
	if input < 0 {
		return "", errors.New("input cannot be negative")
	}

	if input == 0 {
		return "", errors.New("input cannot be zero")
	}

	return fmt.Sprintf("Valid input: %d", input), nil
}
