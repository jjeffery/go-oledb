package oledb

import (
	"fmt"
)

type HResult uint32

var (
	S_OK    = HResult(1)
	S_FALSE = HResult(2)

	E_NOTIMPL     = HResult(0x80004001)
	E_NOINTERFACE = HResult(0x80004002)
	E_POINTER     = HResult(0x80004003)
)

func (hr HResult) Succeeded() bool {
	return (hr & 0x80000000) == 0
}

func (hr HResult) Failed() bool {
	return (hr & 0x80000000) != 0
}

func (hr HResult) String() string {
	switch hr {
	case S_OK:
		return "S_OK"
	case S_FALSE:
		return "S_FALSE"
	case E_NOTIMPL:
		return "E_NOTIMPL"
	case E_NOINTERFACE:
		return "E_NOINTERFACE"
	case E_POINTER:
		return "E_POINTER"
	}

	return fmt.Sprintf("0x%08x", uint32(hr))
}

func HResultFromInt(hr int) HResult {
	ui := uint32(hr)
	return HResult(ui)
}
