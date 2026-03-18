package unwrap

import (
	"errors"
	"fmt"
)

var ErrorNotFound = errors.New("Not Found")

func HttpGet() (any, error) {
	return nil, fmt.Errorf("Operation failed: %w", ErrorNotFound)
}
