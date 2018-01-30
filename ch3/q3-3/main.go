package main

import (
	"archive/zip"
	"os"
	//"strings"
	"strings"
	"io"
)

// zipファイルの書き込み
func main (){
	file, err := os.Create("newfile.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	text, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	sReader := strings.NewReader("Q3.3: zipファイルの書き込み")
	io.Copy(text, sReader)
}

