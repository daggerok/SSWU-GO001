package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond) // work simulation
	}
}

func main() {
	fmt.Println("Starting goroutine...")
	go printNumbers()
	time.Sleep(5 * 500 * time.Millisecond)
	fmt.Println("Goroutine completed")
}
