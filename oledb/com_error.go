package oledb

import (
	"fmt"
)

// ComError contains details about COM errors.
type ComError struct {
	hr         HResult
	methodName string
}

func newComError(hr HResult, methodName string) *ComError {
	return &ComError{
		hr:         hr,
		methodName: methodName,
	}
}

// Code returns the COM error code
func (e *ComError) HResult() HResult {
	return e.hr
}

// MethodName returns the name of the method that returned the error
func (e *ComError) MethodName() string {
	return e.methodName
}

func (e *ComError) String() string {
	return fmt.Sprintf("%s: %s", e.methodName, e.hr.String())
}

// Returns error details as a formatted string
func (e *ComError) Error() string {
	return e.String()
}
