package oledb

import (
	"unsafe"
)

var (
	iid_IDBProperties = &GUID{0x0c733a8a, 0x2a1c, 0x11ce, [8]byte{0xad, 0xe5, 0x00, 0xaa, 0x00, 0x44, 0x77, 0x3d}}
)

type IDBProperties struct {
	vtbl *vtblIDBProperties
}

type vtblIDBProperties struct {
	vtblIUnknown    vtblIUnknown
	getProperties   uintptr
	getPropertyInfo uintptr
	setProperties   uintptr
}

func (dbi *IDBProperties) Unknown() *IUnknown {
	return (*IUnknown)(unsafe.Pointer(dbi))
}

func (dbi *IDBProperties) AddRef() {
	dbi.Unknown().AddRef()
}

func (dbi *IDBProperties) Release() {
	dbi.Unknown().Release()
}

func (dbi *IDBProperties) GetProperties() error {
	method := NewMethod(unsafe.Pointer(dbi), dbi.vtbl.getProperties)
	hr := method.Call0()
	if hr.Failed() {
		return newComError(hr, "IDBProperties.GetProperties")
	}
	return nil
}

func (dbi *IDBProperties) GetPropertyInfo() error {
	method := NewMethod(unsafe.Pointer(dbi), dbi.vtbl.getPropertyInfo)
	hr := method.Call0()
	if hr.Failed() {
		return newComError(hr, "IDBProperties.GetPropertyInfo")
	}
	return nil
}

func (dbi *IDBProperties) SetProperties() error {
	method := NewMethod(unsafe.Pointer(dbi), dbi.vtbl.setProperties)
	hr := method.Call0()
	if hr.Failed() {
		return newComError(hr, "IDBProperties.SetProperties")
	}
	return nil
}
