// https://gist.github.com/nbari/386af0fa667ae03daf3fbc80e3838ab0o

package main

import (
    "fmt"
    "syscall"
)

// select属のシステムコール(macOS: kqueue)によるI/O多重化
// ./testフォルダ内の変更を監視してイベントを待つ
func main() {
    kq, err := syscall.Kqueue()
    if err != nil {
        panic(err)
    }
    // 監視対象のファイルディスクリプタを取得
    fd, err := syscall.Open("./test", syscall.O_RDONLY,   0)
    if err != nil {
        panic(err)
    }

    // 監視したいイベントの構造体を作成
    ev1 := syscall.Kevent_t {
        Ident: uint64(fd),
        Filter: syscall.EVFILT_VNODE,
        Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
        Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE,
        Data:   0,
        Udata:  nil,
    }
    timeout := syscall.Timespec{
        Sec:  0,
        Nsec: 0,
    }

    // イベント待ちの無限ループ
    for {
        // kevent作成
        events := make([]syscall.Kevent_t, 10)
        nev, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, &timeout)
        if err != nil {
            panic(err)
        }

        // イベントを確認
        for i := 0; i < nev; i++ {
            fmt.Printf("Event [%d] -> %+v\n", i, events[i])
        }
    }
}

