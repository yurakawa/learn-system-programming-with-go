package main

import (
    "fmt"
    "sync"
)

// このブロックのコードはinit()をつかうことで
func initialize() {
    fmt.Println("初期化処理")
}
var once sync.Once


func main() {
    // 3回呼び出しても一度しか呼ばれない
    once.Do(initialize)
    once.Do(initialize)
    once.Do(initialize)
}
