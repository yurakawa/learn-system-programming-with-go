package main

import (
    "bufio"
    "fmt"
    "os/exec"
)

// リアルタイムな入出力

// countプログラムを起動して標準出力に(stdout)というプリフィックスを付けつつリアルタイムでリダイレクトするサンプル
// go build -o count count.go を実行語 go run main.go
func main() {
    count := exec.Command("./count")
    stdout, _ := count.StdoutPipe()
    go func() {
        scanner := bufio.NewScanner(stdout)
        for scanner.Scan() {
            fmt.Printf("(stdout) %s\n", scanner.Text())
        }
    }()
    err := count.Run()
    if err != nil {
        panic(err)
    }

}

