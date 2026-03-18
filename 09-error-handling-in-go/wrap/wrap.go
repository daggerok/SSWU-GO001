package wrap

import (
	"errors"
	"fmt"
)

func WrapError() error {
	return fmt.Errorf("Failed to do something: %w", errors.New("Error 404: Not Found"))
}
