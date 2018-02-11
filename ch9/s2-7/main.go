package main

import (
    "os"
    "fmt"
)

func main() {
    if len(os.Args) == 2 {
        fmt.Printf("%s [exec file name]", os.Args[0])
        os.Exit(1)
    }
    info, err := os.Stat(os.Args[1])
    info2, err := os.Stat(os.Args[2])
    if err == os.ErrNotExist {
        fmt.Printf("file not found: %s\n", os.Args[1])
    } else if err != nil {
        panic(err)
    }
    // inode番号やデバイス番号での同一性チェック
    if os.SameFile(info, info2) {
        fmt.Println("同じファイル")
    }
}

