package user32

import (
    "github.com/Viva-Victoria/Ferret/core"
    "github.com/Viva-Victoria/Ferret/windows/winuser"
    "syscall"
)

func CreateWindow(className, windowName string, style winuser.Style, x, y, width, height uint32, parent, menu, instance syscall.Handle) (syscall.Handle, error) {
    classNamePtr, err := core.StringPointer(className)
    if err != nil {
        return 0, err
    }
    
    windowNamePtr, err := core.StringPointer(windowName)
    if err != nil {
        return 0, err
    }
    
    ret, _, err := User32.
        Procedure("CreateWindowExW").
        Call(uintptr(0),
            classNamePtr,
            windowNamePtr,
            uintptr(style),
            uintptr(x),
            uintptr(y),
            uintptr(width),
            uintptr(height),
            uintptr(parent),
            uintptr(menu),
            uintptr(instance),
            uintptr(0))
    
    if ret == 0 {
        return 0, err
    }
    
    return syscall.Handle(ret), nil
}

func DefaultWindowProcedure(window syscall.Handle, message winuser.WindowMessage, w, l uintptr) uintptr {
    ret, _, _ := User32.
        Procedure("DefWindowProcW").
        Call(uintptr(window), uintptr(message), w, l)
    
    return ret
}

func DestroyWindow(window syscall.Handle) error {
    ret, _, err := User32.
        Procedure("DestroyWindow").
        Call(uintptr(window))
    if ret == 0 {
        return err
    }
    
    return nil
}