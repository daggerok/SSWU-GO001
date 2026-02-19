package firstsuccess_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/04-fail-fast-and-first-success-goroutines-channels/internal/firstsuccess"
)

func TestFirstSuccess(t *testing.T) {
	urls := []string{
		"https://www.google.com/",
		"https://www.yahoo.com/",
		"https://timeout.com",
	}

	t.Log(">>> RUNNING: FIRST SUCCESS")
	res, err := firstsuccess.FirstSuccess(urls)

	// Мы ожидаем успех от Google или Yahoo раньше, чем наступит таймаут
	if err != nil {
		t.Fatalf("FirstSuccess should have succeeded, but got error: %v", err)
	}

	if res == "" {
		t.Error("FirstSuccess returned empty result")
	}

	t.Logf("FirstSuccess winner: %s", res)
}
