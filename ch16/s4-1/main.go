package main

import (
    "time"
    "fmt"
    "os"
)

func main() {
    // 5秒
    fmt.Println(5 * time.Second);
    // 10秒
    fmt.Println(10 * time.Second);
    // 10分30秒
    time.ParseDuration("10m30s")

    fmt.Println("---")

    // 現在時刻
    fmt.Println(time.Now())
    // 指定日時を作成
    fmt.Println(time.Date(2017,  time.August, 26, 11, 50, 30, 0, time.Local))
    // ファーマットを指定してパース
    fmt.Println(time.Parse(time.Kitchen, "11:30AM"))
    // Epochタイム から作成
    fmt.Println(time.Unix(1503673200, 0))

    fmt.Println("---")

    // 3時間後の時間
    fmt.Println(time.Now().Add(3 * time.Hour))
    // ファイル変更日時が何日前か知る
    fileInfo, _ := os.Stat("README.md")
    fmt.Printf("%v 前\n", time.Now().Sub(fileInfo.ModTime()))
    // 時間を1時間ごとに丸める
    fmt.Println(time.Now().Round(time.Hour))
}

