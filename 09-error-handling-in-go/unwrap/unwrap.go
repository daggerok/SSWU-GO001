package unwrap

import (
	"errors"
	"fmt"
)

var ErrorNotFound = errors.New("Not Found")

// Use error as Excepted Error, when the failure is a part of normal operation and should be handled: not found,
// input validation. etc...
func HttpGet() (any, error) {
	return nil, fmt.Errorf("Operation failed: %w", ErrorNotFound)
}
