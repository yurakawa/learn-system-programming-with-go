package main

import (
	"io"
	"net"
	"os"
	"net/http"
	"fmt"
	"bufio"
)

// io.Readerを満たす構造体で、よくつかうもの: ネットワーク通信の読み込み [ net.Conn ]
func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	// http.ReadResponse()でhttpのレスポンスをパースする
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダを表示してみる
	fmt.Println(res.Header)
	// ボディを表示してみる。最後にはClose()すること
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
