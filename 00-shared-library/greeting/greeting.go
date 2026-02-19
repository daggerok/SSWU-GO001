package greeting

import (
	"fmt"
	"strings"
)

// Greeting prints a message. In Go, we use pointers (*string)
// to represent a value that could be nil.
func Greeting(name string) string {
	if name == "" || strings.TrimSpace(name) == "" {
		return "Hello, Anonymous!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
