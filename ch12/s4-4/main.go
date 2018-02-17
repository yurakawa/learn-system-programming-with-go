package main

import (
    "fmt"
    "os"
    "strconv"
)

// シグナルを他のプロセスに送る
func main () {
    if len(os.Args) < 2 {
        fmt.Printf("usage: %s [pid]\n", os.Args[0])
        return
    }

    // 第1引数で指定されたプロセスIDを数値に変換
    pid, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }

    process, err := os.FindProcess(pid)
    if err != nil {
        panic(err)
    }

    // シグナルを送る
    process.Signal(os.Kill)
    // Killの場合は次のショートカットも利用可能
    process.Kill()
}

