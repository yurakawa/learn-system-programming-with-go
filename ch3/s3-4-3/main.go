package main

import (
	"io"
	"net"
	"os"
)

// io.Readerを満たす構造体で、よくつかうもの: ネットワーク通信の読み込み [ net.Conn ]
func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}
