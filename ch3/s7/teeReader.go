package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// 読み込まれた内容を別のio.Writerに書き出す。
func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	// データを読み捨てる
	_, _ = ioutil.ReadAll(teeReader)

	// けどバッファに残ってる
	fmt.Println(buffer.String())
}
