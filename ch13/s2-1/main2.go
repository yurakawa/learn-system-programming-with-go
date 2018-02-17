package main

import (
    "fmt"
    "time"
)

// goroutineと情報共有
func main() {
    tasks := []string{
        "cmake ..",
        "cmake . --build Release",
        "cpack",
    }
    for _, task := range tasks {
        go func() {
            // goroutineが起動するときにはループが回りきって
            // 全部のtaskが最後のタスクになってしまう
            fmt.Println(task)
        }()
    }
    time.Sleep(time.Second)
}
