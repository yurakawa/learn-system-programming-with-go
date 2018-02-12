package main

import (
    "fmt"
    "os"
    "path/filepath"
)

// パスを分割する
func main() {
    dir, name := filepath.Split(os.Getenv("GOPATH"))
    fmt.Printf("Dir: %s, Name: %s\n", dir, name)
}

