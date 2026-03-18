package custom

import "fmt"

type CustomError struct {
	Code int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func HttpGet() (any, error) {
	return nil, &CustomError{Code: 404, Message: "Not Found"}
}
