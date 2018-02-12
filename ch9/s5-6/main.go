package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    // boolean
    fmt.Println(filepath.Match("image-*.png", "image-100.png"))

    // array
    files, err := filepath.Glob("./*.png")
    if err != nil {
        panic(err)
    }
    fmt.Println(files)

}
