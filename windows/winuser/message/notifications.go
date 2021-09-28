package message

import (
    . "github.com/Viva-Victoria/Ferret/windows/winuser"
)

const (
    ActivateApp            WindowMessage = 0x001C
    CancelMode             WindowMessage = 0x001F
    ChildActivate          WindowMessage = 0x0022
    Close                  WindowMessage = 0x0010
    Compacting             WindowMessage = 0x0041
    Create                 WindowMessage = 0x0001
    Destroy                WindowMessage = 0x0002
    Enable                 WindowMessage = 0x000A
    EnterSizeMove          WindowMessage = 0x0231
    ExitSizeMove           WindowMessage = 0x0232
    GetIcon                WindowMessage = 0x007F
    GetMinMaxInfo          WindowMessage = 0x0024
    InputLangChange        WindowMessage = 0x0051
    InputLangChangeRequest WindowMessage = 0x0050
)
