package main

import (
    "fmt"
    "os"
)

// ファイルの属性の取得
func main() {
    if len(os.Args) == 1 {
            fmt.Printf("%s [exec file name]", os.Args[0])
            os.Exit(1)
        }
    info, err := os.Stat(os.Args[1])
    if err == os.ErrNotExist {
            fmt.Printf("file not found: %s\n", os.Args[1])
        } else if err != nil {
            panic(err)
        }
    fmt.Println("FileInfo")
    fmt.Printf("  ファイル名: %v\n", info.Name())
    fmt.Printf("  サイズ: %v\n", info.Size())
    fmt.Printf("  変更日時 %v\n", info.ModTime())
    fmt.Println("Mode()")
    fmt.Printf("  ディレクトリ？ %v\n", info.Mode().IsDir())
    fmt.Printf("  読み書き可能な通常ファイル？ %v\n", info.Mode().IsRegular())
    fmt.Printf("  Unixのファイルアクセス権限ビット %o\n", info.Mode().Perm())
    fmt.Printf("  モードのテキスト表現 %v\n", info.Mode().String())
}
