package main

import (
    "os"
    "fmt"
    "path/filepath"
)

// ディレクトリのパスとファイル名とを連結する
func main() {
    fmt.Printf("Temp File Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))
}
