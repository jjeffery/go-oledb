// +build windows

package oledb

import (
	"syscall"
	"unsafe"
)

type method struct {
	target   uintptr
	function uintptr
}

func NewMethod(target uintptr, function uintptr) method {
	return method{
		target:   target,
		function: function,
	}
}

func (m method) Call0() int {
	hr, _, _ = syscall.Syscall(m.function, 1, m.target, 0, 0)
	return hr
}

func (m method) Call1(a1 uintptr) int {
	hr, _, _ = syscall.Syscall(m.function, 2, m.target, a1, 0)
	return hr
}

func (m method) Call2(a1 uintptr, a2 uintptr) int {
	hr, _, _ = syscall.Syscall(m.function, 2, m.target, a1, a2)
	return hr
}
