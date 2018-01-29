package main

import(
	"os"
	"io"
)

// io.Readerを満たす構造体で、よくつかうもの: ファイル入力 [ os.File ]

func main() {
	srcFile, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	// defer で現在のスコープが終了したらファイルをクローズする
	defer srcFile.Close()

	destFile, err := os.Create("wasCopied.go")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		panic(err)
	}
}
