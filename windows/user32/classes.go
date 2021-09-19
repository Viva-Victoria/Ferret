package user32

import (
    "github.com/Viva-Victoria/Ferret/windows/winuser"
    "syscall"
    "unsafe"
)

type WindowClass struct {
    Size            uint32
    Style           uint32
    WindowProcedure uintptr
    ClassExtra      int32
    WindowExtra     int32
    Instance        syscall.Handle
    Icon            syscall.Handle
    Cursor          syscall.Handle
    Background      winuser.Color
    MenuName        *uint16
    ClassName       *uint16
    IconSmall       syscall.Handle
}

func (w *WindowClass) Pointer() uintptr {
    return uintptr(unsafe.Pointer(w))
}

func (w *WindowClass) Register() (uint16, error) {
    ret, _, err := User32.
        Procedure("RegisterClassExW").
        Call(w.Pointer())
    
    if ret == 0 {
        return 0, err
    }
    
    return uint16(ret), nil
}
