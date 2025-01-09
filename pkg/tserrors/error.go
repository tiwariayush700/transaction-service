package tserrors

import "fmt"

// Error defines a custom error type with a code and message
type Error struct {
	Code    int
	Message string
}

// Error implements the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// New creates a new Error
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
