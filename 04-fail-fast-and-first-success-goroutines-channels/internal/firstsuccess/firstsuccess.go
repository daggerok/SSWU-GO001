package firstsuccess

import (
	"context"
	"fmt"
	"time"

	"github.com/daggerok/SSWU-GO001/04-fail-fast-and-first-success-goroutines-channels/internal/fetch"
)

// FirstSuccess returns first success result and fails with error or timout otherwise
func FirstSuccess(urls []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results := make(chan string, len(urls))
	errors := make(chan error, len(urls))

	for _, url := range urls {
		go func() {
			res, err := fetch.Fetch("FirstSuccess", ctx, url)
			if err != nil {
				errors <- err
				return
			}

			select {
			case results <- res:
				cancel()
			case <-ctx.Done():
			}
		}()
	}

	errorCount := 0
	for {
		select {
		case res := <-results:
			return res, nil
		case err := <-errors:
			errorCount++
			fmt.Printf("logged error: %v\n", err)
			if errorCount == len(urls) {
				return "", fmt.Errorf("all failed")
			}
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}
