package oledb

import (
	"strings"
	"testing"
)

func TestNewComError(t *testing.T) {
	var testData = []struct {
		hr            HResult
		methodName    string
		expectedError string
	}{
		{HResultFromInt(-1), "methodName", "methodName: 0xFFFFFFFF"},
		{HResult(0x8007000E), "IUnknown::QueryInterface", "IUnknown::QueryInterface: 0x8007000E: Not enough storage"},
		{E_POINTER, "IUnknown::QueryInterface", "IUnknown::QueryInterface: E_POINTER: Invalid pointer"},
		{E_NOINTERFACE, "IUnknown::QueryInterface", "IUnknown::QueryInterface: E_NOINTERFACE: No such interface"},
	}

	for _, td := range testData {
		err := newComError(td.hr, td.methodName)
		if hr := err.HResult(); hr != td.hr {
			t.Errorf("Code() failed, expected %s, actual = %s", td.hr.String(), hr.String())
		}
		if methodName := err.MethodName(); methodName != td.methodName {
			t.Errorf("MethodName() failed, expected %s, actual = %s", td.methodName, methodName)
		}
		if errorText := err.Error(); !strings.HasPrefix(errorText, td.expectedError) {
			t.Errorf("Error() failed, expected to start with '%s\n', actual = '%s'", td.expectedError, errorText)
		}
	}
}
