package core

import (
    "sync"
    "syscall"
)

type Library struct {
    name string
    dll *syscall.LazyDLL
    cache sync.Map
}

func (l *Library) Procedure(name string) *syscall.LazyProc {
    if proc, ok := l.cache.Load(name); ok {
        return proc.(*syscall.LazyProc)
    }
    
    proc := l.dll.NewProc(name)
    if proc != nil {
        l.cache.Store(name, proc)
    }
    return proc
}

func NewLibrary(name string) *Library {
    return &Library{
        name: name,
        dll: syscall.NewLazyDLL(name),
    }
}