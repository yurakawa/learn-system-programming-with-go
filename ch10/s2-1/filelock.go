// +build darwin dragonfly freebsd linux netbsd openbsd

// syscall.Flock()によるPOSIX系OSでのファイルロック
package main

import (
    "sync"
    "syscall"

)

type FileLock struct {
    l sync.Mutex
    fd int
}

func NewFileLock(filename string) * FileLock {
    if filename == "" {
        panic("filename needed")
    }
    fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY,0750)
    if err != nil {
        panic(err)
    }
    return &FileLock{fd: fd}
}

func (m *FileLock) Lock() {
    m.l.Lock()
    if err := syscall.Flock(m.fd, syscall.LOCK_EX); err != nil {
        panic(err)
    }
}

func (m *FileLock) Unlock() {
    if err := syscall.Flock(m.fd, syscall.LOCK_UN); err != nil {
        panic(err)
    }
    m.l.Unlock()
}
