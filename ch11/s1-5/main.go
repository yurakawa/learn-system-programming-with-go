package main

import (
    "fmt"
    "os"
)

// 作業フォルダ
func main () {
    wd, _ := os.Getwd()
    fmt.Println(wd)
}

