package main

// io.Reader /io.Writerでストリームを自由に操る

import (
	"bytes"
	"io"
	"os"
)

// 引数で渡されたio.Readerのすべての入出力をつなげる
func main() {
	header := bytes.NewBufferString("---- HEADER ----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("---- FOOTER ----\n")

	reader := io.MultiReader(header, content, footer)
	// すべての reader をつなげた出力が表示
	io.Copy(os.Stdout, reader)
}
