package main

import (
    "fmt"
    "net"
    "path/filepath"
    "os"
)

// データグラム型のUnixドメインソケット(クライアント)
func main() {
    conn, err := net.Dial("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
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

