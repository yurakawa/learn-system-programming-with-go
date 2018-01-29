package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1行目
2行目
3行目`

// テキスト解析用のio.Reader関連機能
// bufio.Reader を使って改行区切りで読み込む
func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n",line)
		if err == io.EOF {
			break
		}
	}
}
