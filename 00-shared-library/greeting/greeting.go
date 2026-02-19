package greeting

import "fmt"

// Greeting prints a message. In Go, we use pointers (*string)
// to represent a value that could be nil.
func Greeting(name *string) string {
	if name == nil || *name == "" {
		return "Hello, Anonymous!"
	} else {
		return fmt.Sprintf("Hello, %s!", *name)
	}
}
