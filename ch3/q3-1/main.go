package main

import (
	"os"
	"io"
)

func main() {
	file, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, _ := os.Create("new.txt")
	defer newFile.Close()
	io.Copy(newFile, file)
}

