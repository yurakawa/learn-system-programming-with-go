package main

import (
    "fmt"
    "os"
)

// プロセスID
func main() {
    fmt.Printf("プロセスID: %d\n", os.Getpid())
    fmt.Printf("親プロセスID: %d\n", os.Getppid())
}
