package show

import (
    . "github.com/Viva-Victoria/Ferret/windows/winuser"
)

const (
    Hide            ShowWindowFlag = 0
    ShowNormal      ShowWindowFlag = 1
    ShowMinimized   ShowWindowFlag = 2
    ShowMaximized   ShowWindowFlag = 3
    NoActivate      ShowWindowFlag = 4
    Show            ShowWindowFlag = 5
    Minimize        ShowWindowFlag = 6
    ShowMinNoActive ShowWindowFlag = 7
    ShowNotActive   ShowWindowFlag = 8
    Restore         ShowWindowFlag = 9
    ShowDefault     ShowWindowFlag = 10
    ForceMinimize   ShowWindowFlag = 11
)
