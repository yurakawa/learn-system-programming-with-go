package main

import (
    "fmt"
    "path/filepath"
)

// パスのクリーン化
func main() {
    // パスをそのままクリーンにする
    fmt.Println(filepath.Clean("./path/filepath/../path.go")) // path/path.go

    // パスを絶対パスに整形
    abspath, _ := filepath.Abs("path/filepath/path_unix.go")
    fmt.Println(abspath)

    // パスを相対パスに整形
    relpath, _ := filepath.Rel("/usr/local/go/src",
                                "/usr/local/go/src/path/filepath/path.go")
    fmt.Println(relpath) // path/filepath/path.go
}

