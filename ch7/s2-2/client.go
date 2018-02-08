package main

import (
    "fmt"
    "net"
)

// UDPで接続するクライアント側の実装例
func main() {

    conn, err := net.Dial("udp4", "localhost:8888")
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    fmt.Println("Sending to server")
    _, err = conn.Write([]byte("Hello from Client"))
    if err != nil {
        panic(err)
    }
    fmt.Println("Receiveing from server")
    buffer := make([]byte, 1500)
    length, err := conn.Read(buffer)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Received: %s\n", string(buffer[:length]))
}
