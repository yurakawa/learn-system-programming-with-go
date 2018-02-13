package main

import (
    "fmt"
    "time"
)

// main.goのロックを取得して10秒後に解放する
// exec: go run main.go filelock.go
func main() {
    l := NewFileLock("main.go")
    fmt.Println("try locking")
    l.Lock()
    fmt.Println("locked!")
    time.Sleep(10 * time.Second)
    l.Unlock()
    fmt.Println("unlock")

}