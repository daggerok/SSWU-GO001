package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond) // work simulation
	}
}

func printAsyncMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // notify on func scope exit as a last step
	fmt.Println(msg)
}

func main() {
	fmt.Println("Starting goroutine...")
	go printNumbers()
	time.Sleep(5 * 500 * time.Millisecond)
	fmt.Println("Goroutine completed")

	// WaitGroup
	var wg sync.WaitGroup
	wg.Add(3)
	fmt.Println("Waiting for 3 goroutines to finish")
	go printAsyncMessage("Async", &wg)
	go printAsyncMessage("hello", &wg)
	go printAsyncMessage("message", &wg)
	wg.Wait() // block util all goroutines have finished
	fmt.Println("Goroutines completed")

	// Unbuffered Channels
	unbuffered := make(chan string)
	go func() {
		message := "Hello from goroutine"
		fmt.Println("Sending:", message)
		unbuffered <- message // send; will be blocked if no receiver ready to take the value
	}()

	message := <-unbuffered // receive; will be blocked if no sender or channel isn't closed properly
	fmt.Println("Received:", message)

	// Buffered Channels
	// When processing data in stages, unbuffered channels ensure that
	// the next stage does not start until the previous stage completes.
	bufferedQueue := make(chan string, 2)
	bufferedQueue <- "First message"
	bufferedQueue <- "2nd message"
	fmt.Println("Received:", <-bufferedQueue)
	fmt.Println("Received:", <-bufferedQueue)

	// Closed Channel
	// Sender side:
	ch := make(chan string, 2)
	ch <- "1st message"
	ch <- "2nd message"
	close(ch) // only sender should close
	// Receiver side:
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	// Receiving from a closed channel
	msg, ok := <-ch
	fmt.Printf("Received: '%s' %v\n", msg, ok)

	// Select Statement for Multiple Channels
	// define delayed sender infrastructure:
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "ch1 message"
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "ch2 message"
	}()
	// defined delayed receiver infrastructure:
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:fmt.Println("msg1 received:", msg1)
		case msg2 := <-ch2:fmt.Println("msg2 received:", msg2)
		case <- time.After(501 * time.Millisecond):fmt.Println("Timeout")
		}
	}
}
