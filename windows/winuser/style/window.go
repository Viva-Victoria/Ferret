package style

import (
    . "github.com/Viva-Victoria/Ferret/windows/winuser"
)

const (
    Border           Style = 0x00800000
    Caption          Style = 0x00C00000
    Child            Style = 0x40000000
    ClipChildren     Style = 0x02000000
    ClipSiblings     Style = 0x04000000
    Disabled         Style = 0x08000000
    DialogFrame      Style = 0x00400000
    HorizontalScroll Style = 0x00100000
    Iconic           Style = 0x20000000
    Maximize         Style = 0x01000000
    MaximizeBox      Style = 0x00010000
    Minimize         Style = 0x20000000
    MinimizeBox      Style = 0x00020000
    Overlapped       Style = 0x00000000
    OverlappedWindow       = Overlapped | Caption | SystemMenu | ThickFrame | MinimizeBox | MaximizeBox
    Popup            Style = 0x80000000
    PopupWindow            = Popup | Border | SystemMenu
    SizeBox          Style = 0x00040000
    SystemMenu       Style = 0x00080000
    TabStop          Style = 0x00010000
    ThickFrame       Style = 0x00040000
    Visible          Style = 0x10000000
    VerticalScroll   Style = 0x00200000
)
