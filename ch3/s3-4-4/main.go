package main

import (
	"bytes"
	"strings"
)

// bytes.Buffer の 初期化コード（コンパイル不通)
func main() {

	// io.Readerを満たす構造体で、よくつかうもの: メモリに備えた内容をio.Readerとして読み出すバッファ [ bytes.Buffer, bytes.Reader, strings.Reader]
	// 空のバッファ(実態なので&buffer1のようにポインタ値を渡す必要がある)
	var buffer1 bytes.Buffer
	// バイト列で初期化
	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	// 文字列で初期化
	buffer3 := bytes.NewBufferString("初期文字列")


	// bytes.Reader は bytes.NewReaderで作成
	bReader1 := bytes.NewReader([]byte{0x10, 0x20, 0x30})
	bReader2 := bytes.NewReader([]byte("文字列をバイト配列にキャストして設定"))

	// strings.Readerはstrings.NewReader()関数で作成
	sReader := strings.NewReader("Readerの出力内容は文字列で渡す")

}
