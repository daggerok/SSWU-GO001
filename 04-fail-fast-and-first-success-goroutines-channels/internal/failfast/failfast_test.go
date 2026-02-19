package failfast_test

import (
	"strings"
	"testing"

	"github.com/daggerok/SSWU-GO001/04-fail-fast-and-first-success-goroutines-channels/internal/failfast"
)

func TestFailFast(t *testing.T) {
	urls := []string{
		"https://www.google.com/",
		"https://www.yahoo.com/",
		"https://timeout.com",
	}

	t.Log(">>> RUNNING: FAIL FAST")
	res, err := failfast.FailFast(urls)

	// Мы ожидаем ошибку, так как timeout.com должен "уронить" группу
	if err == nil {
		t.Errorf("FailFast should have failed due to timeout.com, but got result: %s", res)
	} else {
		t.Logf("FailFast correctly failed with error: %v", err)
		errStr := err.Error()
		// Проверяем, что ошибка НЕ является ни одной из допустимых
		isKnownError := strings.Contains(errStr, "timeout.com") ||
			strings.Contains(errStr, "context canceled") ||
			strings.Contains(errStr, "429") ||
			strings.Contains(errStr, "too many requests")

		if !isKnownError {
			t.Errorf("Unexpected error message: '%v' doesn't contain timeout.com or context canceled", err)
		}
	}
}
