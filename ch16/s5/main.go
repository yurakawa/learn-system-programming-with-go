package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now.Format(time.RFC822))
    // 25 Feb 18 23:31 JST

    fmt.Println(now.Format("2006/01/02 03:04:05 MST"))
    // 2018/02/25 11:31:47 JST
}
