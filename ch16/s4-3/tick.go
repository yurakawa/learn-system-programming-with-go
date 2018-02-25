package main

import(
    "fmt"
    "time"
)

// 等間隔で通知するtime.Tick()
func main() {
    fmt.Println("waiting 5 seconds")
    for now := range time.Tick(5 * time.Second) {
        fmt.Println("now:", now)
    }
}