package main

import (
	"fmt"

	"github.com/daggerok/SSWU-GO001/00-shared-library/greeting"
	"github.com/daggerok/SSWU-GO001/01-hello-gopher/internal/uppercaser"
)

func main() {
	upperName := uppercaser.ToUpper("go")
	greet := greeting.Greeting(upperName)
	fmt.Println(greet)
}
