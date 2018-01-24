package main

import (
	"bytes"
	"fmt"
)

// 書かれた内容をためておいてまとめて受け取れるbytes.Buffer
func main() {
	var buffer bytes.Buffer

	buffer.Write([]byte("bytes.Buffer example1\n"))
	fmt.Println(buffer.String())

	buffer.Write([]byte("bytes.Buffer example2\n"))
	fmt.Println(buffer.String())
}
