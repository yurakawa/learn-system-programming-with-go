package main

import (
    "fmt"
    "os"
    "path/filepath"
)

// whichコマンド
// PATH環境変数のパス一覧を取得してきて、それをfilepath.SplitList()でここのパスに分割
// その後、各パスの下に最初の引数で指定されてあ実行ファイルがあるかをチェック
func main() {
    if  len(os.Args) == 1 {
        fmt.Printf("%s [exec file name]", os.Args[0])
        os.Exit(1)
    }
    // filepath.SplitList()は複数のパスを一つにまとめたものを分解する
    for _, path := range filepath.SplitList(os.Getenv("PATH")) {
        execpath := filepath.Join(path, os.Args[1])
        _, err := os.Stat(execpath)
        if !os.IsNotExist(err) {
            fmt.Println(execpath)
            return
        }
    }
    os.Exit(1)
}
