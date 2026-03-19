package main

import (
	"github.com/daggerok/SSWU-GO001/11-debugging-techniques/log"
)

func main() {
	_, _ = log.Divide(10, 0)
	_, _ = log.Divide(10, 1)
	_, _ = log.Divide(10, 2)
}
