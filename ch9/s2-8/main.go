package main

import (
    "os"
    "time"
)

func main() {
    filename := "setting.txt"

    // ファイルのモードを変更
    os.Chmod(filename, 0644)

    // ファイルのオーナーを変更
    os.Chown(filename, os.Getuid(), os.Getgid())

    // ファイルの最終アクセス日時と変更日時を変更
    os.Chtimes(filename, time.Now(), time.Now())
}
