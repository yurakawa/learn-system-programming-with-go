package main

import (
    "os"
    "fmt"
    "io"
)

func open(filename string) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    io.WriteString(file, "New file contentt\n")
}

// リンク
func main() {
    oldName := "oldfile.txt"
    open(oldName)

    // ハードリンク
    os.Link(oldName, "newfile.txt")

    // シンボリックリンク
    os.Symlink(oldName, "newfile-symlink.txt")

    // シンボリックリンクのリンク先を取得
    link, err := os.Readlink("newfile-symlink.txt")
    if err != nil {
        panic(err)
    }
    fmt.Printf("link to: %v", link)
}
