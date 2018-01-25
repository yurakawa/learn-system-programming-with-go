package main

import (
	"os"
	"io"
	"compress/gzip"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")

	// gzip圧縮(書き込まれたデータをgzip圧縮して、予め渡されていたos.Fileに中継する)
	file2, err2 := os.Create("test.txt.gz")
	if err != nil {
		panic(err2)
	}
	writer2 := gzip.NewWriter(file2)
	writer2.Header.Name = "text.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer2.Close()
}
