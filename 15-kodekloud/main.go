package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Printf("Name reading error: %T %v\n", err, err)
		return
	}

	fmt.Print("Enter your age: ")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Printf("Age reading error: %T %v\n", err, err)
		return
	}

	fmt.Printf("Hello, %q and you are %d years old.\n", name, age)
}
