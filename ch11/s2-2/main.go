package main

import (
    "fmt"
    "os"
)

// 環境変数
func main() {
    fmt.Println(os.Environ())
    // 環境変数が埋め込まれた文字列を展開
    fmt.Println(os.ExpandEnv("${HOME}/gobin"))
}
