package user32

import (
    "github.com/Viva-Victoria/Ferret/core"
    "github.com/Viva-Victoria/Ferret/windows/winuser"
    "syscall"
)

func LoadCursor(cursor winuser.CursorIdentifier) (syscall.Handle, error) {
    ret, _, err := User32.
        Procedure("LoadCursorW").
        Call(uintptr(0), cursor.Pointer())
    if ret == 0 {
        return 0, err
    }
    
    return syscall.Handle(ret), nil
}

func LoadCursorFromFile(file string) (syscall.Handle, error) {
    fileNamePtr, err := core.StringPointer(file)
    if err != nil {
        return 0, err
    }
    
    ret, _, err := User32.
        Procedure("LoadCursorFromFileW").
        Call(uintptr(0), fileNamePtr)
    if ret == 0 {
        return 0, err
    }
    
    return syscall.Handle(ret), nil
}
