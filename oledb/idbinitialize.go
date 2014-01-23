package oledb

import (
	"unsafe"
)

var (
	iid_IDBInitialize = &GUID{0x0c733a8b, 0x2a1c, 0x11ce, [8]byte{0xad, 0xe5, 0x00, 0xaa, 0x00, 0x44, 0x77, 0x3d}}
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

func (dbi *IDBInitialize) QueryDBProperties() (*IDBProperties, error) {
	if ptr, err := dbi.Unknown().QueryInterface(iid_IDBProperties); err != nil {
		return nil, err
	} else {
		return (*IDBProperties)(ptr), nil
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
	if hr.Failed() {
		return newComError(hr, "IDBInitialize.Initialize")
	}
	return nil
}

func (dbi *IDBInitialize) Uninitialize() error {
	method := NewMethod(unsafe.Pointer(dbi), dbi.vtbl.uninitialize)
	hr := method.Call0()
	if hr.Failed() {
		return newComError(hr, "IDBInitialize.Uninitialize")
	}
	return nil
}

func CreateDBInitialize(progId string) (*IDBInitialize, error) {
	ptr, err := createInstance(progId, iid_IDBInitialize)
	if err != nil {
		return nil, err
	}

	return (*IDBInitialize)(ptr), nil
}
