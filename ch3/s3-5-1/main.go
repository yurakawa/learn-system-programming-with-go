package main

import (
	"io"
	"strings"
	"os"
	"fmt"
)

// バイナリ解析用のio.Reader 関連機能
// 必要な部分を切り出す io.LimitReader / io.SectionReader
func main() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)

	fmt.Print("\n")

	lReader := io.LimitReader(reader, 16)
	io.Copy(os.Stdout, lReader)
}
