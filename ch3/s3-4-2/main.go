package main

import(
	"os"
	"io"
)

// io.Readerを満たす構造体で、よくつかうもの: ファイル入力 [ os.File ]

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	// defer で現在のスコープが終了したらファイルをクローズする
	defer file.Close()
	io.Copy(os.Stdout, file)

}
