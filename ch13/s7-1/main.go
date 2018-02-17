package main

import (
    "fmt"
    "sync"
)

// IDをインクリメントするコードが同時に一つしか時刻されないようにする

var id int

func generateId(mutex *sync.Mutex) int {
    mutex.Lock()
    defer mutex.Unlock()
    id++
    return id
}
func main() {
    // sync.Mutex構造体の変数宣言
    // 次の宣言をしてもポインタ型になるだけで正常に動作する
    // mutex := new(sync.Mutex)
    var  mutex sync.Mutex

    for i := 0; i < 100; i++ {
        go func() {
            fmt.Printf("id: %d\n", generateId(&mutex))
        }()
    }
}

