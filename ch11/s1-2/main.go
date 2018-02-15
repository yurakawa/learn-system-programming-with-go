package main

import (
    "fmt"
    "os"
    "syscall"
)

// プロセスグループ・セッショングループ
func main() {
    sid, _ := syscall.Getsid(os.Getpid())
    fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n",
                                syscall.Getpgrp(), sid)
}
