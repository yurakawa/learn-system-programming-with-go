package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目 2行目 3行目`

// テキスト解析用のio.Reader関連機能
// bufio.Scannerを使う(デリミターは削除される)
func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	// 単語区切りに変更
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		fmt.Printf("%#v\n",scanner.Text())
	}
}
