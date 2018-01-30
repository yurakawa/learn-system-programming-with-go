package main

import (
	"archive/zip"
	"os"
	//"strings"
)

// zipファイルの書き込み
func main (){
	file, err := os.Create("newfile.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	_, err = zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
}
