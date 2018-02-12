package main

import (
    "os"
    "fmt"
    "os/user"
    "path/filepath"
)

// ~ も観葉変数も展開した上でパスをクリーン化する
func Clean2(path string) string {
    if len(path) > 1 && path[0:2] == "~/" {
        my, err := user.Current()
        if err != nil {
            panic(err)
        }
        path = my.HomeDir + path[1:]
    }
    path = os.ExpandEnv(path)
    return filepath.Clean(path)
}

func main() {
    path := os.ExpandEnv("${GOPATH}/src/github.com")
    fmt.Println(Clean2(path))
    path = os.ExpandEnv("~/.bashrc")
    fmt.Println(Clean2(path))
}