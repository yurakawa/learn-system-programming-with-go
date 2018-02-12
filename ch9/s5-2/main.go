package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {
    path := "/folder1/folder2/example.txt"

    // パスを分割する
    dir, name := filepath.Split(path)
    fmt.Printf("Dir: %s, Name: %s\n", dir, name)

    // パスをセパレート文字化する
    fragments := strings.Split(path, string(filepath.Separator))
    fmt.Println(fragments)

    // パスの最後の要素を返す
    fmt.Println(filepath.Base(path))

    // パスのディレクトリ部を返す
    fmt.Println(filepath.Dir(path))

    // ファイルの拡張子を返す
    fmt.Println(filepath.Ext(path))
}

