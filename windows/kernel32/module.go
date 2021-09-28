package kernel32

import (
    "syscall"
)

func GetModuleHandle() (syscall.Handle, error) {
    ret, _, err := Kernel32.Procedure("GetModuleHandleW").Call(uintptr(0))
    if ret == 0 {
        return 0, err
    }
    
    return syscall.Handle(ret), nil
}
