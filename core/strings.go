package core

import (
    "syscall"
    "unsafe"
)

func StringPointer(s string) (uintptr, error) {
    ptr, err := syscall.UTF16PtrFromString(s)
    if err != nil {
        return 0, err
    }
    
    return uintptr(unsafe.Pointer(ptr)), nil
}
