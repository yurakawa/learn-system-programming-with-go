package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	// Writeが受け取るのはバイト列なので変換してから渡している
	file.Write([]byte("os.File examplen"))
	file.Close()
}
