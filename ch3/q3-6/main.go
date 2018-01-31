package main

import(
	"strings"
	"io"
	"os"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system	    = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

// 次の3つの文字列を3つの入力ストリーム(io.Reader)とし、下記に示すmain()
// 関数のコメント部にコードを追加して、最後のio.Copy()で「ASCII」の
// 文字列が出力されるようにしてみてください

func main() {
	var stream io.Reader

	// ここにio パッケージをつかったコードを書く
	a := io.NewSectionReader(programming,5, 1)
	s := io.NewSectionReader(system, 0, 1)
	c := io.NewSectionReader(computer, 0, 1)
	i := io.NewSectionReader(programming, 8, 1)
	i2 := io.NewSectionReader(programming, 8, 1)
	stream = io.MultiReader(a,s,c,i,i2)
	//

	io.Copy(os.Stdout,stream)
}
