// +build windows

package oledb

import (
	"syscall"
	"unsafe"
)

type method struct {
	target   unsafe.Pointer
	function uintptr
}

func NewMethod(target unsafe.Pointer, function uintptr) method {
	return method{
		target:   target,
		function: function,
	}
}

func (m method) Call0() HResult {
	hr, _, _ := syscall.Syscall(m.function, 1, uintptr(m.target), 0, 0)
	return HResult(hr)
}

func (m method) Call1(a1 uintptr) HResult {
	hr, _, _ := syscall.Syscall(m.function, 2, uintptr(m.target), a1, 0)
	return HResult(hr)
}

func (m method) Call2(a1 uintptr, a2 uintptr) HResult {
	hr, _, _ := syscall.Syscall(m.function, 2, uintptr(m.target), a1, a2)
	return HResult(hr)
}
