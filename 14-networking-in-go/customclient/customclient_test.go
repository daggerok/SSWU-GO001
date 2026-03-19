package customclient_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/daggerok/SSWU-GO001/14-networking-in-go/customclient"
)

func TestCustomHttpClient(t *testing.T) {
	t.Run("should send HTTP request by using custom client", func(t *testing.T) {
		customHttpClient := customclient.NewHTTPClient()
		resp, err := customHttpClient.Get("https://api.github.com")
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
