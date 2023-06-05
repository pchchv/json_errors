package json_errors

// BaseError interface reveals additional information about the error.
// Implements a built-in error interface.
type BaseError interface {
	Error() string
}

// baseError is a simple error struct.
type baseError struct {
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}
