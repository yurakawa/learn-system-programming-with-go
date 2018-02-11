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

// 追記モード
func append() {
    file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    io.WriteString(file, "Appened content\n")
}

func main() {
    open()
    read()
    append()
    read()
}
