package panic

import "fmt"

// DivideAndPanic use panic in case of Critical Errors, where program cannot continue execution: out of memory,
// unrecoverable I/O failures, etc...
func DivideAndPanic(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

// DividePanicAndRecover use panic as Programming Errors, where panic is appropriate for scenarios that represents code
// bugs: index out of bounds, divide by zero, etc...
func DividePanicAndRecover(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("divide by zero")
	}

	return a / b
}
