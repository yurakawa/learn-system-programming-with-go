package main

import (
    // "math/rand"でも動く //
    "crypto/rand"
    "encoding/hex"
    "fmt"
)

func main() {
    a := make([]byte, 20)
    rand.Read(a)
    fmt.Println(hex.EncodeToString(a))
}
