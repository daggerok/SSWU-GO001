package customserver

import (
	"fmt"
	"net/http"
	"time"
)

func NewHTTPServer() *http.Server {
	server := &http.Server{
		Addr:    ":8080",
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, "Message received")
		}),
	}
	return server
}
