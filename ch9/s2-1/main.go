package main

import (
    "fmt"
    "io"
    "os"
)

//ファイル作成/読み込み

// 新規作成
func open() {
    file, err := os.Create("textfile.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    io.WriteString(file, "New file contentt\n")
}

// 読み込み
func read() {
    file, err := os.Open("textfile.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    fmt.Println("Read file:")
    io.Copy(os.Stdout, file)
}

func main() {
    open()
    read()
    read()
}
