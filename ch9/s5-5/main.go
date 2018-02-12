package main

import (
    "os"
    "fmt"
    "os/user"
)

func main() {
    // 環境変数絵を展開
    path := os.ExpandEnv("${GOPATH}/src/github.com")
    fmt.Println(path)

    // ホームディレクトリを取得
    my, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Println(my.HomeDir)


}