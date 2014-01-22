package oledb

import (
	"fmt"
)

const (
	S_OK = 0

	E_NOTIMPL     = 0x80004001
	E_NOINTERFACE = 0x80004002
	E_POINTER     = 0x80004003
)

var (
	mapHresultToName map[int]string
)

func init() {
	mapHresultToName = make(map[int]string)
	mapHresultToName[E_NOINTERFACE] = "E_NOINTERFACE"
	mapHresultToName[E_POINTER] = "E_POINTER"
}

// ComError contains details about COM errors.
type ComError struct {
	hr         int
	methodName string
}

func newComError(hr int, methodName string) *ComError {
	return &ComError{
		hr:         hr,
		methodName: methodName,
	}
}

// Code returns the COM error code
func (e *ComError) Code() int {
	return e.hr
}

// MethodName returns the name of the method that returned the error
func (e *ComError) MethodName() string {
	return e.methodName
}

// Returns error details as a formatted string
func (e *ComError) Error() string {
	if name := mapHresultToName[e.hr]; name != "" {
		return fmt.Sprintf("%s: hr = %s", e.methodName, name)
	}
	return fmt.Sprintf("%s: hr = 0x%08x", e.methodName, uint32(e.hr))
}
