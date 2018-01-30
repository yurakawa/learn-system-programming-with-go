package main

import (
	"crypto/rand"
	"io"
	"os"
)

// テスト用の適当なサイズ(1024バイト)のファイルを作成
func main() {
	file, err := os.Create("random.txt")
	if err != nil {
		panic(err)
	}
	_, err = io.CopyN(file, rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
}


