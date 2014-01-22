package oledb

import (
	"unsafe"
)

var (
	IID_IDBInitialize = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
)

type IDBInitialize struct {
	vtbl *vtblIDBInitialize
}

type vtblIDBInitialize struct {
	vtblIUnknown vtblIUnknown
	initialize   uintptr
	uninitialize uintptr
}

func (dbi *IDBInitialize) Unknown() *IUnknown {
	return (*IUnknown)(unsafe.Pointer(dbi))
}

func (unk *IUnknown) QueryDBInitialize() (*IDBInitialize, error) {
	if ptr, err := unk.QueryInterface(IID_IDBInitialize); err != nil {
		return nil, err
	} else {
		return (*IDBInitialize)(ptr), nil
	}
}

func (dbi *IDBInitialize) AddRef() {
	dbi.Unknown().AddRef()
}

func (dbi *IDBInitialize) Release() {
	dbi.Unknown().Release()
}

func (dbi *IDBInitialize) Initialize() error {
	method := NewMethod(unsafe.Pointer(dbi), dbi.vtbl.initialize)
	hr := method.Call0()
	if hr != S_OK {
		return newComError(hr, "IDBInitialize.Initialize")
	}
	return nil
}

func (dbi *IDBInitialize) Uninitialize() error {
	method := NewMethod(unsafe.Pointer(dbi), dbi.vtbl.uninitialize)
	hr := method.Call0()
	if hr != S_OK {
		return newComError(hr, "IDBInitialize.Uninitialize")
	}
	return nil
}
