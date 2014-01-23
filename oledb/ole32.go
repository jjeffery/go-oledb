package oledb

import (
	"strings"
	"syscall"
	"unsafe"
)

const (
	clsctx_INPROC_SERVER   = 1
	clsctx_INPROC_HANDLER  = 2
	clsctx_LOCAL_SERVER    = 4
	clsctx_INPROC_SERVER16 = 8
	clsctx_REMOTE_SERVER   = 16
	clsctx_ALL             = clsctx_INPROC_SERVER | clsctx_INPROC_HANDLER | clsctx_LOCAL_SERVER
	clsctx_INPROC          = clsctx_INPROC_SERVER | clsctx_INPROC_HANDLER
	clsctx_SERVER          = clsctx_INPROC_SERVER | clsctx_LOCAL_SERVER | clsctx_REMOTE_SERVER
)

var (
	ole32Dll, _             = syscall.LoadDLL("ole32.dll")
	procCoInitialize, _     = ole32Dll.FindProc("CoInitialize")
	procCoCreateInstance, _ = ole32Dll.FindProc("CoCreateInstance")
	procCLSIDFromProgID, _  = ole32Dll.FindProc("CLSIDFromProgID")
	procCLSIDFromString, _  = ole32Dll.FindProc("CLSIDFromString")
)

func Initialize() error {
	code, _, _ := procCoInitialize.Call(uintptr(0))
	hr := HResult(code)

	if hr.Failed() {
		return hr.ComError("CoInitialize")
	}
	return nil
}

func coCreateInstance(clsid *GUID, iid *GUID) (unsafe.Pointer, error) {
	if iid == nil {
		iid = IID_IUnknown
	}

	var unk unsafe.Pointer

	ret, _, _ := procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(clsid)),
		0,
		clsctx_SERVER,
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&unk)))

	hr := HResult(ret)
	if hr.Failed() {
		return nil, hr.ComError("CoCreateInstance")
	}

	return unk, nil
}

func clsidFromProgId(progId string) (*GUID, error) {
	var clsid GUID
	lpszProgId := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(progId)))
	ret, _, _ := procCLSIDFromProgID.Call(lpszProgId, uintptr(unsafe.Pointer(&clsid)))
	hr := HResult(ret)
	if hr.Failed() {
		return nil, hr.ComError("CLSIDFromProgID")
	}
	return &clsid, nil
}

func clsidFromString(s string) (*GUID, error) {
	var clsid GUID
	lpsz := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
	ret, _, _ := procCLSIDFromString.Call(lpsz, uintptr(unsafe.Pointer(&clsid)))
	hr := HResult(ret)
	if hr.Failed() {
		return nil, hr.ComError("CLSIDFromString")
	}

	return &clsid, nil
}

func createInstance(progId string, iid *GUID) (unsafe.Pointer, error) {
	var clsid *GUID
	var err error

	if strings.HasPrefix(progId, "{") {
		clsid, err = clsidFromString(progId)
	} else {
		clsid, err = clsidFromProgId(progId)
	}

	if err != nil {
		return nil, err
	}

	unk, err := coCreateInstance(clsid, iid)
	if err != nil {
		return nil, err
	}

	return unk, nil
}
