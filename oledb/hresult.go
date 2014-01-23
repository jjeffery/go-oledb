package oledb

import (
	"fmt"
	"syscall"
	"unicode/utf16"
)

type HResult uint32

var (
	S_OK    = HResult(1)
	S_FALSE = HResult(2)

	E_NOTIMPL     = HResult(0x80004001)
	E_NOINTERFACE = HResult(0x80004002)
	E_POINTER     = HResult(0x80004003)
	E_ABORT       = HResult(0x80004004)
	E_FAIL        = HResult(0x80004005)

	E_OUTOFMEMORY = HResult(0x8007000E)
	E_INVALIDARG  = HResult(0x80070057)

	CO_E_NOTINITIALIZED = HResult(0x800401F0)
)

func (hr HResult) Succeeded() bool {
	return (hr & 0x80000000) == 0
}

func (hr HResult) Failed() bool {
	return (hr & 0x80000000) != 0
}

func (hr HResult) String() string {

	var name string

	switch hr {
	case S_OK:
		name = "S_OK"
	case S_FALSE:
		name = "S_FALSE"
	case E_NOTIMPL:
		name = "E_NOTIMPL"
	case E_NOINTERFACE:
		name = "E_NOINTERFACE"
	case E_POINTER:
		name = "E_POINTER"
	case CO_E_NOTINITIALIZED:
		name = "CO_E_NOTINITIALIZED"
	default:
		name = fmt.Sprintf("0x%08X", uint32(hr))
	}

	var flags uint32 = syscall.FORMAT_MESSAGE_FROM_SYSTEM | syscall.FORMAT_MESSAGE_ARGUMENT_ARRAY | syscall.FORMAT_MESSAGE_IGNORE_INSERTS
	buf := make([]uint16, 320)
	n, err := syscall.FormatMessage(flags, 0, uint32(hr), 0, buf, nil)
	if n > 0 && err == nil {

		// trim trailing newline
		for n > 0 && (buf[n-1] == '\n' || buf[n-1] == '\r') {
			n--
		}

		msg := string(utf16.Decode(buf[:n]))
		return fmt.Sprintf("%s: %s", name, msg)
	}

	return name
}

func (hr HResult) Error() string {
	return hr.String()
}

func (hr HResult) ComError(methodName string) *ComError {
	return newComError(hr, methodName)
}

func HResultFromInt(hr int) HResult {
	ui := uint32(hr)
	return HResult(ui)
}
