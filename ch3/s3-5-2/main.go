package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// バイナリ解析用のio.Reader 関連機能
// エンディアン変換
func main() {
	// 32bitのビッグエンディアンのデータ(10000)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// エンディアンの変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
