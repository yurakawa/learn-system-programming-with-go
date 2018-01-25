package main

import (
	"io"
	"os"
	"net"
)

func main() {
	// net.Dial()で返されるnet.Connインタフェースはio.Writerとio.Readerとのハイブリッドのインタフェース
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
