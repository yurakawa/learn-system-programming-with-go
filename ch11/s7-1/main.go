package main

import (
    "fmt"
    "github.com/mattn/go-colorable" // エスケープシーケンスの環境差を吸収する
    "github.com/mattn/go-isatty" // ターミナルか判別する
    "io"
    "os"
)

var data = "\033[34m\033[47m\033[4mB\033[31me\n\033[24m\033[30mOS\033[49m\033[m\n"

// go-colorableとgo-isattyを利用して接続先がターミナルのときはエスケープシーケンスを表示し
// そうでないときはエスケープし～kネスを場外するフィルターを使い分ける
func main() {
    var stdOut io.Writer
    if isatty.IsTerminal(os.Stdout.Fd()){
        stdOut = colorable.NewColorableStdout()
    } else {
        stdOut = colorable.NewNonColorable(os.Stdout)
    }
    fmt.Fprintln(stdOut, data)
}
