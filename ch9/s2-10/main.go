package main

import (
    "fmt"
    "os"
)

// ディレクトリ情報の取得
func main() {
    dir, err := os.Open("/")
    if err != nil {
        panic(err)
    }
    // Readdir()はos.FileInfoの配列を返す
    // 正の整数を与えると、最大でその個数の要素だけを返す
    // 0以下の数値を渡すと、ディレクトリ内の全要素を返す
    fileInfos, err := dir.Readdir(-1)
    if err != nil {
        panic(err)
    }
    for _, fileInfo := range fileInfos {
        if fileInfo.IsDir() {
            fmt.Printf("[Dir] %s\n", fileInfo.Name())
        } else {
            fmt.Printf("[File] %s\n", fileInfo.Name())
        }
    }
}
