package divide

import "errors"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}
