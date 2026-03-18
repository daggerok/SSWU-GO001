package main

import (
	"errors"
	"fmt"
)

//type error interface {
//	Error() string
//}

func DoSomething() (string, error) {
	return "", errors.New("Something failed")
}

func main() {
	res, err := DoSomething()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(res)
}
