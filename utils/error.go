package utils

import "fmt"

// CustomError wraps errors with additional context
type CustomError struct {
    Context string
    Err     error
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("%s: %v", e.Context, e.Err)
}

// Define specific error types
type APIError struct {
    StatusCode int
    Message    string
}

func (e *APIError) Error() string {
    return fmt.Sprintf("API Error (Status %d): %s", e.StatusCode, e.Message)
}

// WrapError creates a new CustomError with context
func WrapError(context string, err error) error {
    return &CustomError{
        Context: context,
        Err:     err,
    }
}