package main

import (
    "fmt"
    "time"
)

// ワンショットで1回だけ待つtime.After()
func main() {
    fmt.Println("waiting 5 seconds")
    after := time.After(5 * time.Second)
    <- after
    fmt.Println("done")
}
