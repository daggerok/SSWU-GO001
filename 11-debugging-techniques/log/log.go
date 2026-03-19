package log

import (
	"log"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, log.Output(2, "0 is not a divide") // log error message
	}
	result := a / b
	log.Printf("Debug: %d / %d = %d", a, b, result) // log debug message
	return result, nil
}
