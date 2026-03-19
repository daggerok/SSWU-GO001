package client_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	t.Run("should send HTTP request", func(t *testing.T) {
		resp, err := http.Get("https://api.github.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(body))
	})
}
