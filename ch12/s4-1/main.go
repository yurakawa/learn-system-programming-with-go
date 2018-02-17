package main

import (
    "fmt"
    "os/signal"
    "syscall"
    "time"
)

// シグナルを無視する
func main() {
    // 最初の10秒はCtrl + Cで止まる
    fmt.Println("Accept Ctrl + C for 10 second")
    time.Sleep(time.Second * 10)

    // 可変長引数で任意の数のシグナルを設定可能
    signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

    // 次の10秒はCtrl + C を無視する
    fmt.Println("Ignore Ctrl + C for 10 second")
    time.Sleep(time.Second * 10)

}


