package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"time"
)

// functions

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

// methods

type Rectangle struct {
	Width, Height float64
}

// value receivers

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// pointer receiver

//goland:noinspection GoMixedReceiverTypes
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// interfaces, inheritance and polymorphism

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

//goland:noinspection GoMixedReceiverTypes
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

//type Square struct {
//	Width float64
//}
//
//func (s Square) Area() float64 {
//	return s.Width * s.Width
//}
//
//func (r Square) Perimeter() float64 {
//	return 2 * (r.Width + r.Width)
//}
//
////goland:noinspection GoMixedReceiverTypes
//func (r *Square) Scale(factor float64) {
//	r.Width *= factor
//}
type Square struct {
	Rectangle
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func printAnything(anything interface{}) {
	fmt.Println(anything)
}

func printAny(any any) {
	fmt.Println(any)
}

func printIfString(any any) {
	str, ok := any.(string)
	if ok {
		fmt.Println(str)
	}
}

func printByType(any any) {
	switch v := any.(type) { // type is available only inside switch operator
	case int:fmt.Println("int:", v)
	case float64:fmt.Println("float64:", v)
	case string:fmt.Println("string:", v)
	default:fmt.Println("unknown type:", reflect.TypeOf(any), v)
	}
}

// structs

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person Person
	Position string
}

// json

type PersonJson struct {
	Name string	`json:"name"`
	Age  int	`json:"age"`
}

type EmployeeJson struct {
	PersonJson PersonJson	`json:"personJson"`
	Position string			`json:"position"`
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

	// arrays
	var arr [5]int // declare and initialize array with defaults zeroes
	fmt.Println("arr:", arr)

	array := [3]string{"one", "two", "three"}
	fmt.Println("array:", array) // declare and initialize array with specific strings

	// slices
	sliceLiteral := []int{1, 2, 3}
	fmt.Println("sliceLiteral:", sliceLiteral)

	slice := make([]int, 4) // creates a slice of len 4 with zeroes values
	fmt.Println("slice:", slice)

	subSlice := array[0:1]
	fmt.Println("subSlice:", subSlice) // create a slice from `array` of strings from index 0 (inclusive) to 1 (exclusive)

	// methods
	rectangle := Rectangle{10, 5}
	fmt.Println("rectangle:", rectangle, "area:", rectangle.Area(), "perimeter:", rectangle.Perimeter())

	rectangle.Scale(2)
	fmt.Println("scaled rectangle by 2:", rectangle, "area:", rectangle.Area(), "perimeter:", rectangle.Perimeter())

	// interfaces
	circle := Circle{7}
	fmt.Print("circle: ", circle, " area: "); printArea(circle)
	circle.Scale(2)
	fmt.Print("scaled circle by 2: ", circle, " area: "); printArea(circle)
	printAnything(circle)
	printAny(circle)
	printIfString(circle)
	printIfString("circle")

	printByType(1234)
	printByType(math.Pi)
	printByType(circle)
	printByType("circle")

	square := Square{Rectangle{4, 4}}
	fmt.Print("square: ", square, " perimeter: ", square.Perimeter(), " area: "); printArea(square)
	square.Scale(3)
	fmt.Println("scaled square by 3:", square, "perimeter:", square.Perimeter(), "area:", square.Area())
	printByType(square)

	// structs
	employee := Employee{
		Person: Person{"Max", 42},
		Position: "Principle Software Engineer",
	}
	fmt.Println("employee:", employee.Person.Name, "(", employee.Position, ")")

	var emptyEmployee Employee // = Employee{}
	fmt.Println("empty employee:", emptyEmployee)

	// json
	ej, err := json.Marshal(EmployeeJson{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("string(ej):", string(ej))

	// anonymous struct
	anonymousPerson := struct {
		Name string
		Age  int
	}{
		Name: "Max",
		Age: 42,
	}
	fmt.Println("anonymousPerson:", anonymousPerson)
}
