package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// optimizing CPU-bound code

	// use all available CPUs cores:
	runtime.GOMAXPROCS(runtime.NumCPU())

	// run goroutines
	go expensiveTask(100000000)
	go expensiveTask(100000000)
	go expensiveTask(100000000)

	//// wait (for user manual input):
	// fmt.Scanln()
	time.Sleep(1 * time.Second)

	// optimizing memory-bound code
	// - avoid unnecessary memory allocations: inside the loops and frequently executed code
	// - reuse memory: use sync.Pool to reuse memory between operations, especially for temporary allocations
	// - batching operations: group operations together to reduce memory churn and reduce GC load
	var bufferPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024) // creates a new buffer
		},
	}

	// acquire buffer:
	buf := bufferPool.Get().([]byte)
	fmt.Println("Buffer size:", len(buf))

	//nolint:staticcheck // SA6002: сознательно передаем слайс, а не указатель // return buffer to hte pool:
	bufferPool.Put(buf)

	anotherBuf := bufferPool.Get().([]byte)
	fmt.Println("Reused buffer with a size:", len(anotherBuf))
}

func expensiveTask(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}
