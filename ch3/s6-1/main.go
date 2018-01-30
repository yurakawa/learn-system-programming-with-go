package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目
2行目
3行目`

// テキスト解析用のio.Reader関連機能
// bufio.Scannerを使う(デリミターは削除される)
func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan(){
		fmt.Printf("%#v\n",scanner.Text())
	}
}
