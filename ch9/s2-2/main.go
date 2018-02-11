package main

import (
    "os"
)

func main() {
    // フォルダを1階層だけ作成
    os.Mkdir("setting", 0755)

    // 深いフォルダを1回で作成
    os.MkdirAll("setting/myapp/networksettings", 0755)
}
