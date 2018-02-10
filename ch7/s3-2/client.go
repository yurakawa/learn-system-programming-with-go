package main

import (
    "fmt"
    "net"
)


/**
 * UDPのマルチキャストの実装例(クライント側)
 */

 // ソケットを開いて、サーバが10秒に1回送信するパケットを待って表示する
func main() {
    fmt.Println("Listen tick server at 224.0.0.1:9999")
    address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
    if err != nil {
        panic(err)
    }
    listener, err := net.ListenMulticastUDP("udp", nil, address)
    defer listener.Close()

    buffer := make([]byte, 1500)

    for {
        length, remoteAddress, err := listener.ReadFromUDP(buffer)
        if err != nil {
            panic(err)
        }
        fmt.Printf("Server %v\n", remoteAddress)
        fmt.Printf("Now    %s\n", string(buffer[:length]))
    }
}
