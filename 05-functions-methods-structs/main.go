package main

import (
	"fmt"
	"time"
)

func swap(a, b int) (y, z int) {
	y = b
	z = a
	//return y, z
	return
}

func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero: %w", err)
	}
	return a / b, nil
}

func divideAndPanic(a, b float64) float64 {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

func divideWithoutPanic(a, b float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	return divideAndPanic(a, b)
}

func sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func multiplier(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func trackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func main() {
	defer trackTime(time.Now(), "main")

	a, b := 1, 2
	y, z := swap(a, b)
	fmt.Printf("swapped %d & %d are %d & %d\n", a, b, y, z)

	res, err := divide(0, 0)
	fmt.Printf("expected divide by zero error %t and zero result %.2f\n", err, res)

	sumOf123 := sum(1, 2, 3)
	fmt.Println("sum of 1 + 2 + 3 =", sumOf123)

	sliceOf1234567890 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sumOf1234567890 := sum(sliceOf1234567890...)
	fmt.Println("sum of 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 =", sumOf1234567890)

	//closures
	x := 10
	doubleX := func() int {
		return x * 2 //captures x from parent scope
	}
	fmt.Printf("x is %d doubleX closure execution is %d\n", x, doubleX())

	twoTimes := multiplier(2)
	fmt.Printf("twoTimes(3) = %d\n", twoTimes(3))

	//map operations
	ages := map[string]float32{
		"Max": 41.8,
		"Elena": 33.6,
		"Ameliia": 9.5,
		"Martin": 1.8,
	}
	name := "Martin"
	if age, exists := ages[name]; exists {
		fmt.Printf("%s is %.1f years old\n", name, age)
	} else {
		fmt.Printf("%s's age not found\n", name)
	}

	fmt.Printf("divideAndPanic(1, 2) %.1f\n", divideAndPanic(1, 2))
	//fmt.Printf("divideAndPanic(1, 0) %.1f\n", divideAndPanic(1, 0))
	//panic: divide by zero
	//
	//goroutine 1 [running]:
	//main.divideAndPanic(...)
	///path/to/SSWU-GO001/05-functions-methods-structs/main.go:24
	//main.main()
	///path/to/SSWU-GO001/05-functions-methods-structs/main.go:90 +0x51c
	//exit status 2

	fmt.Printf("divideWithoutPanic(1, 0) %.1f\n", divideWithoutPanic(1, 0))
}
