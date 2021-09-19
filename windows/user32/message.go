package user32

import (
    "github.com/Viva-Victoria/Ferret/windows/winuser/message"
    "syscall"
)

func DispatchMessage(message *message.Message) {
    _, _, _ = User32.
        Procedure("DispatchMessageW").
        Call(message.Pointer())
}

func GetMessage(message *message.Message, handle syscall.Handle, min, max uint32) (bool, error) {
    ret, _, err := User32.
        Procedure("GetMessageW").
        Call(message.Pointer(), uintptr(handle), uintptr(min), uintptr(max))
    
    retInt32 := int32(ret)
    if retInt32 == -1 {
        return false, err
    }
    
    return retInt32 != 0, nil
}

func PostQuitMessage(exitCode int32) {
    _, _, _ = User32.Procedure("PostQuitMessage").Call(uintptr(exitCode))
}

func TranslateMessage(message *message.Message) {
    _, _, _ = User32.Procedure("TranslateMessage").Call(message.Pointer())
}