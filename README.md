go-oledb
========

OLEDB for Go language

I'm not going to continue with this: it makes more sense to use ADO, and mattn has already
done this: see https://github.com/mattn/go-adodb.

I might keep the code for a while, because I quite like the implementation for some of the lower-level code (eg hresult.go, method_windows.go, iunknown.go), and I'll keep it in case I want to do some COM-interop for non-scriptable interfaces in future.
