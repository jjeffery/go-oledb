package oledb

import (
	"syscall"
)

var (
	ole32, _        = syscall.LoadDLL("ole32.dll")
	coInitialize, _ = ole32.FindProc("CoInitialize")
)

func init() {
	code, _, _ := coInitialize.Call(uintptr(0))
	hr := HResult(code)

	if hr.Failed() {
		panic(newComError(hr, "CoInitialize").String())
	}
}
