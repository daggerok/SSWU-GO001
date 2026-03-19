package customserver_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/daggerok/SSWU-GO001/14-networking-in-go/customclient"
	"github.com/daggerok/SSWU-GO001/14-networking-in-go/customserver"
)

func TestNewHTTPServer(t *testing.T) {
	t.Run("should create a new http server", func(t *testing.T) {
		server := customserver.NewHTTPServer()
		if server == nil {
			t.Error("server is nil")
		}
		defer server.Close()

		go func() {
			fmt.Println("Server is running on port 8080...")
			//goland:noinspection GoUnhandledErrorResult
			server.ListenAndServe()
		}()

		client := customclient.NewHTTPClient()
		if client == nil {
			t.Error("client is nil")
		}

		resp, err := client.Get("http://localhost:8080")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println("Body:", string(body))
	})
}
