package main

import (
    "fmt"
    "net"
)

// UDPの接続を受け付けるサーバ側の実装例
func main() {
    fmt.Println("Server is runnning at localhost:8888")
    conn, err := net.ListenPacket("udp", "localhost:8888")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    buffer := make([]byte, 1500)
    for {
        length, remoteAddress, err := conn.ReadFrom(buffer)
        if err != nil {
            panic(err)
        }
        fmt.Printf("Received from %v: %v\n",
            remoteAddress, string(buffer[:length]))
        _, err = conn.WriteTo([]byte("Hello from Server"),remoteAddress)
        if err != nil {
            panic(err)
        }
    }
}
