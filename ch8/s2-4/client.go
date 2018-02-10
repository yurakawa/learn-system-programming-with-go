package main

import (
    "fmt"
    "net"
    "path/filepath"
    "os"
    "log"
)

// データグラム型のUnixドメインソケット(クライアント)
func main() {
    clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-server")
    os.Remove(clientPath)
    // net.Dial()で開いたソケットは一方的な送信用でアドレス(ソケットファイルのパスと結び付けられていないので、返信を受けられないため ListenPacket を使う必要がある
    conn, err := net.ListenPacket("unixgram", clientPath)
    if err != nil {
        panic(err)
    }
    // 送信先のアドレス
    unixServerAddr, err := net.ResolveUnixAddr("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
    var serverAddr net.Addr =  unixServerAddr
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    log.Println("Sending to server")

    _, err = conn.WriteTo([]byte("Hello from Client"), serverAddr)
    if err != nil {
        panic(err)
    }
    fmt.Println("Receiveing from server")
    buffer := make([]byte, 1500)
    length,  _, err := conn.ReadFrom(buffer)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Received: %s\n", string(buffer[:length]))
}

