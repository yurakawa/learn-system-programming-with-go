package main

import (
	"fmt"
	"io"
	"os"
)

// io.Readerを満たす構造体で、よくつかうもの
// 標準入力に対応するオブジェクト os.Stdin.Read()
func main() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input'%s'\n", size, string(buffer))

	}
}

