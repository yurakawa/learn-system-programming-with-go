package main

import (
    "fmt"
    "time"

)

// 1秒に1つずつ数値を出力するプログラム
func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
        time.Sleep(time.Second)
    }
}
