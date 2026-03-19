package main

import "fmt"

type LargeDataSet struct {
	data [10000]int
}

func byValue(largeDataSet LargeDataSet) {
	fmt.Println("Value received:", largeDataSet.data[0])
}

func byPointer(largeDataSet *LargeDataSet) {
	fmt.Println("Pointer received:", largeDataSet.data[0])
}

func main() {
	// Proper Use of Pointers to Avoid Unnecessary Memory Allocations
	ls := LargeDataSet{}
	byValue(ls) // Inefficiently creates a copy of the LargeDataSet
	byPointer(&ls) // Efficiently passes a pointer to the LargeDataSet

	// Managing Slice Capacity and Growth to Avoid Performance Penalties
	slice := make([]int, 0, 100) // slice with size: 0 and capacity: 100
	fmt.Println("slice:", slice)

	// slice with no initial capacity:
	var numbers []int
	for i := 0; i < 1000; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println("Capacity of numbers::", cap(numbers)) // capacity growths automatically more than needed

	// slice with predefined capacity:
	// its recommended to preallocate capacity if size is known to avoid multiple reallocations
	efficientNumbers := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		efficientNumbers = append(efficientNumbers, i)
	}
	// cap helps to check slice capacity and adjust it as needed. use slice effectively for large datasets or
	// collections
	fmt.Println("Capacity of efficient numbers:", cap(efficientNumbers)) // capacity remains steady
}
