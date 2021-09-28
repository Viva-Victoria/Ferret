package message

import (
    "github.com/Viva-Victoria/Ferret/windows/windef"
    "github.com/Viva-Victoria/Ferret/windows/winuser"
    "syscall"
    "unsafe"
)

type Message struct {
    Handle  syscall.Handle
    Message winuser.WindowMessage
    W       uintptr
    L       uintptr
    Time    uint32
    Point   windef.Point
}

func (m *Message) Pointer() uintptr {
    return uintptr(unsafe.Pointer(m))
}
