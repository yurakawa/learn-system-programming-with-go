package main

import (
    "os"
    "io"
)

func open(filename string) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    text := `
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et
dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex
ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu 
fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt
mollit anim id est laborum.
`
    io.WriteString(file, text)

}

func main() {
    filename := "server.log"
    open(filename)

    ////
    // ファイルの削除

    // ファイルや空のディレクトリの削除
    os.Remove(filename)

    os.MkdirAll("workdir/myapp/", 0755)
    // ディレクトリを中身ごと削除
    os.RemoveAll("workdir")

    ////
    // 特定の長さでファイルを切り取る

    // 先頭100バイトで切る
    open(filename)
    os.Truncate(filename, 100)

    // Truncateメソッドを利用する場合
    file, _ := os.Open("server.log")
    file.Truncate(100)

    ////
    // ファイルを移動する/リネームする

    // リネーム
    open(filename)
    os.Rename(filename, "renamed.txt")

    // 移動
    open(filename)
    os.Mkdir("newdir", 0755)
    os.Rename(filename, "newdir/file.txt")
    // os.Rename(filename, "newdir/") // エラーが発生する


}
