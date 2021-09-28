package main

import (
    "github.com/Viva-Victoria/Ferret/windows/kernel32"
    "github.com/Viva-Victoria/Ferret/windows/user32"
    "github.com/Viva-Victoria/Ferret/windows/winuser"
    "github.com/Viva-Victoria/Ferret/windows/winuser/color"
    "github.com/Viva-Victoria/Ferret/windows/winuser/cursor"
    "github.com/Viva-Victoria/Ferret/windows/winuser/message"
    "github.com/Viva-Victoria/Ferret/windows/winuser/style"
    "log"
    "os"
    "os/signal"
    "syscall"
    "unsafe"
)

const ClassName = "MyTestClass"
const UseDefault = 0x80000000

func main() {
    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, os.Kill)
    
    go func() {
        instance, err := kernel32.GetModuleHandle()
        if err != nil {
            panic(err)
        }
    
        cursor, err := user32.LoadCursor(cursor.Arrow)
        if err != nil {
            panic(err)
        }
    
        fn := func(window syscall.Handle, msg winuser.WindowMessage, w, l uintptr) uintptr {
            switch msg {
            case message.Close:
                _ = user32.DestroyWindow(window)
                sig <- nil
            case message.Destroy:
                user32.PostQuitMessage(0)
                sig <- nil
            default:
                return user32.DefaultWindowProcedure(window, msg, w, l)
            }
        
            return 0
        }
    
        classNamePtr, err := syscall.UTF16PtrFromString(ClassName)
        if err != nil {
            panic(err)
        }
    
        class := user32.WindowClass{
            WindowProcedure: syscall.NewCallback(fn),
            Instance: instance,
            Cursor: cursor,
            Background: color.Window + 1,
            ClassName: classNamePtr,
        }
        class.Size = uint32(unsafe.Sizeof(class))
    
        if _, err := class.Register(); err != nil {
            panic(err)
        }
    
        _, err = user32.CreateWindow(
            ClassName,
            "Test",
            style.OverlappedWindow | style.Visible,
            UseDefault,
            UseDefault,
            UseDefault,
            UseDefault,
            0,
            0,
            instance)
        if err != nil {
            panic(err)
        }
        
        for {
            msg := message.Message{}
            gotMessage, err := user32.GetMessage(&msg, 0, 0, 0)
            if err != nil {
                log.Println(err)
                return
            }
        
            if gotMessage {
                user32.TranslateMessage(&msg)
                user32.DispatchMessage(&msg)
            } else {
                break
            }
        }
    }()
    
    <-sig
}
