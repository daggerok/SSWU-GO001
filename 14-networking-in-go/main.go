package main

import (
	"fmt"
	"net/http"
)

func main() {
	// run web server:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Message received")
	})

	// handle requests:
	fmt.Println("Server is running on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
