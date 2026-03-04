package main

import (
	"fmt"
	"time"
)

const Pi float32 = 3.14

func main() {
	fmt.Println("Pi constant:", Pi)

	//name := "Max"
	//age := 42
	//weight := 118.5

	//var name = "Max"
	//var age = 42
	//var weight = 118.5

	//var name string = "Max"
	//var age int8 = 42
	//var weight float32 = 118.5

	//var name, age, weight = "Max", 42, 118.5

	name, age, weight := "Max", 42, 118.5
	fmt.Printf("My name is %s, I'm %d years old and my weight is %.2f.\n", name, age, weight)

	//var pointer *string = &name
	var pointer = &name
	fmt.Println(name, "->", pointer)

	*pointer = "Maksimko"
	fmt.Println(name, "->", pointer)

	_ = "unsued variable workaround with underscore (blank identifier)"

	if //goland:noinspection GoBoolExpressions
	age < 10 {
		fmt.Println(name, "is a minor")
	} else if //goland:noinspection GoBoolExpressions
	age > 65 {
		fmt.Println(name, "is a senior")
	} else {
		fmt.Println(name, "is an adult")
	}

	today := time.Now().Weekday()

	//switch today {
	//case time.Monday:fallthrough
	//case time.Tuesday:fallthrough
	//case time.Wednesday:fallthrough
	//case time.Thursday:fallthrough
	//case time.Friday: fmt.Println("Today is workday")
	//case time.Saturday:fallthrough
	//case time.Sunday: fmt.Println("Today is weekend")
	//}

	switch today {
		case time.Saturday, time.Sunday:
			fmt.Println("Today is weekend")
		default:
			fmt.Println("Today is workday")
	}

	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}

	index := 0
	for index < 3 {
		fmt.Println("index = ", index)
		index++
	}

	numbers := []int8 {0, 1, 2}
	for index, num := range numbers {
		fmt.Printf("i %d num =%d\n", index, num)
	}

	pairs := map[string]int8 {"one": 1, "two": 2, "three": 3}
	for key, value := range pairs {
		fmt.Println(key, "->", value)
	}

	greeting := "Hello, world!"
	for index, value := range greeting {
		fmt.Printf("%#U starts at byte position: %d\n", value, index)
	}

	//infinite loop
	for {
		fmt.Println("I was here")
		break
	}
}
