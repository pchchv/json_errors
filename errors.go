package json_errors

// BaseError interface reveals additional information about the error.
// Implements a built-in error interface.
type BaseError interface {
	Error() string
}
