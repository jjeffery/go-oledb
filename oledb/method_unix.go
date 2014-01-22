// +build !windows

package oledb

import (
	"unsafe"
)

type method struct {
}

func NewMethod(target unsafe.Pointer, function uintptr) method {
	return method{}
}

func (m method) Call0() int {
	return E_NOTIMPL
}

func (m method) Call1(a1 uintptr) int {
	return E_NOTIMPL
}

func (m method) Call2(a1 uintptr, a2 uintptr) int {
	return E_NOTIMPL
}
