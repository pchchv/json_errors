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

func (e *baseError) Error() string {
	msg := "{"

	if e.Message != "" {
		msg += `"message":"` + e.Message + `"`
	}

	if e.Details != "" {
		if e.Details[0] == '{' {
			msg += `,"details":` + e.Details
		} else {
			msg += `,"details":"` + e.Details + `"`
		}
	}

	msg += "}"

	return msg
}
