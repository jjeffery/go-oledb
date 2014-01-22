package oledb

import (
	"testing"
)

func TestNewComError(t *testing.T) {
	var testData = []struct {
		hr            int
		methodName    string
		expectedError string
	}{
		{-1, "methodName", "methodName: hr = 0xffffffff"},
		{0x80000005, "IUnknown::QueryInterface", "IUnknown::QueryInterface: hr = 0x80000005"},
		{E_POINTER, "IUnknown::QueryInterface", "IUnknown::QueryInterface: hr = E_POINTER"},
		{E_NOINTERFACE, "IUnknown::QueryInterface", "IUnknown::QueryInterface: hr = E_NOINTERFACE"},
	}

	for _, td := range testData {
		err := newComError(td.hr, td.methodName)
		if code := err.Code(); code != td.hr {
			t.Errorf("Code() failed, expected %d, actual = %d", td.hr, code)
		}
		if methodName := err.MethodName(); methodName != td.methodName {
			t.Errorf("MethodName() failed, expected %s, actual = %s", td.methodName, methodName)
		}
		if errorText := err.Error(); errorText != td.expectedError {
			t.Errorf("Error() failed, expected '%s', actual = '%s'", td.expectedError, errorText)
		}
	}
}
