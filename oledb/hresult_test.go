package oledb_test

import (
	"github.com/jjeffery/go-oledb/oledb"
	"testing"
)

func TestErrorCodes(t *testing.T) {

	var errorCodes = []oledb.HResult{
		oledb.E_NOTIMPL,
		oledb.E_NOINTERFACE,
		oledb.E_POINTER,
		//oledb.HResult(-400),
	}

	for _, errorCode := range errorCodes {
		if errorCode.Succeeded() {
			t.Errorf("Expected error code not to succeed: %d", int(errorCode))
		}

		if !errorCode.Failed() {
			t.Errorf("Expected error code to fail: %d", int(errorCode))
		}
	}
}

func TestSuccessCodes(t *testing.T) {
	var successCodes = []oledb.HResult{
		oledb.S_OK,
		oledb.S_FALSE,
		oledb.HResult(400),
	}

	for _, successCode := range successCodes {
		if !successCode.Succeeded() {
			t.Errorf("Expected code to succeed: %d", int(successCode))
		}

		if successCode.Failed() {
			t.Errorf("Expected code not to fail: %d", int(successCode))
		}
	}
}
