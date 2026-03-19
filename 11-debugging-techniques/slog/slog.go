package slog

import (
	"fmt"
	"log/slog"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		slog.Error("division error", "operation", "divide", "a", a, "b", b)
		return 0, fmt.Errorf("division by zero")
	}
	result := a / b
	slog.Info("division success", "a", a, "b", b, "result", result) // log debug message
	return result, nil
}
