package fetch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Fetch fetches URL and returns result or error
func Fetch(id string, ctx context.Context, url string) (string, error) {
	if url == "https://timeout.com" {
		select {
		case <-time.After(5 * time.Second):
			return id, fmt.Errorf("%s timeout 5 seconds reached for %s", id, url)
		case <-ctx.Done():
			return id, ctx.Err()
		}
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return id, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) Go-Test-Agent")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return id, err
	}

	if resp == nil {
		return id, fmt.Errorf("%s received nil response", id)
	}

	defer resp.Body.Close()
	if statusCode := resp.StatusCode; statusCode >= http.StatusBadRequest {
		return id, fmt.Errorf("%s unexpected status code: %d for %s", id, statusCode, url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return id, err
	}

	return fmt.Sprintf("%s [%s] OK: len(%d)", id, url, len(body)), nil
}
