package oledb

import (
	"unsafe"
)

var (
	IID_IUnknown = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
)

type IUnknown struct {
	vtbl *vtblIUnknown
}

type vtblIUnknown struct {
	queryInterface uintptr
	addRef         uintptr
	release        uintptr
}

func (unk *IUnknown) QueryInterface(iid *GUID) (ptr unsafe.Pointer, err error) {
	method := NewMethod(unsafe.Pointer(unk), unk.vtbl.queryInterface)
	hr := method.Call2(uintptr(unsafe.Pointer(iid)), uintptr(unsafe.Pointer(&ptr)))

	if hr.Failed() {
		err = newComError(hr, "IUnknown.QueryInterface")
	} else {
		err = nil
	}

	return
}

// QueryUnknown returns a pointer to the base *IUnknown pointer for this
// COM object. Querying for the IUnknown interface is useful when trying
// to verify if two COM pointers refer to the same COM object.
func (unk *IUnknown) QueryIUnknown() (*IUnknown, error) {
	if ptr, err := unk.QueryInterface(IID_IUnknown); err != nil {
		return nil, err
	} else {
		return (*IUnknown)(ptr), nil
	}
}

func (unk *IUnknown) AddRef() {
	method := NewMethod(unsafe.Pointer(unk), unk.vtbl.addRef)
	hr := method.Call0()
	if hr.Failed() {
		panic(newComError(hr, "IUnknown.AddRef").String())
	}
}

func (unk *IUnknown) Release() {
	method := NewMethod(unsafe.Pointer(unk), unk.vtbl.release)
	hr := method.Call0()
	if hr.Failed() {
		panic(newComError(hr, "IUnknown.Release").String())
	}
}
