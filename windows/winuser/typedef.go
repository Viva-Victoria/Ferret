package winuser

import (
    "syscall"
)

type Style uint32

type WindowMessage uint32

type CursorIdentifier uint32

func (c CursorIdentifier) Pointer() uintptr {
    return uintptr(uint16(c))
}

type Color syscall.Handle

type ShowWindowFlag int