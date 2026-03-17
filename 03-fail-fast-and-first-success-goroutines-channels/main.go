package main

import (
	"fmt"

	"github.com/daggerok/SSWU-GO001/03-fail-fast-and-first-success-goroutines-channels/internal/failfast"
	"github.com/daggerok/SSWU-GO001/03-fail-fast-and-first-success-goroutines-channels/internal/firstsuccess"
)

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://www.yahoo.com/",
		"https://timeout.com",
	}

	fmt.Println(">>> RUNNING: FAIL FAST")
	if res, err := failfast.FailFast(urls); err != nil {
		fmt.Printf("FAIL FAST RESULT: Interrupted by error: %v\n", err)
	} else {
		fmt.Printf("FAIL FAST WINNER (should never happen): %s\n", res)
	}

	fmt.Println(">>> RUNNING: FIRST SUCCESS")
	if res, err := firstsuccess.FirstSuccess(urls); err == nil {
		fmt.Println("FIRST SUCCESS WINNER:", res)
	} else {
		fmt.Println("FIRST SUCCESS RESULT:", err)
	}
}
