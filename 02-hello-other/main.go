package main

import (
	"fmt"

	"github.com/daggerok/SSWU-GO001/00-shared-library/greeting"
)

func main() {
	greet := greeting.Greeting("")
	fmt.Println(greet)
}
