package main

import (
    "bufio"
    "fmt"
    "net"
    "net/http"
    "net/http/httputil"
    "strings"
)

// Keep-Alive対応のHTTPクライアント
func main() {
    sendMessages := []string{
        "ASCII",
        "PROGRAMMING",
        "PLUS",
    }
    current := 0
    var conn net.Conn = nil
    // リトライ用にループで全体を囲う
    for {
        var err error
        // まだコネクションを張ってない / エラーでリトライ
        if conn == nil {
            // Dialから行ってconnを初期化
            conn, err = net.Dial("tcp", "localhost:8888")
            if err !=nil {
                panic(err)
            }
            fmt.Printf("Access: %d\n", current)
        }
        // POSTで文字列を送るリクエストを作成
        request, err := http.NewRequest(
            "POST",
            "http://localhost:8888",
            strings.NewReader(sendMessages[current]))
        if err != nil {
            panic(err)
        }
        err = request.Write(conn)
        if err != nil {
            panic(err)
        }

        // サーバーから読み込む。タイムアウトはここでエラーになるのでリトライ
        response, err := http.ReadResponse(bufio.NewReader(conn), request)
        if err != nil {
            fmt.Println("Retry")
            conn = nil
            continue
        }
        // 結果を表示
        dump, err := httputil.DumpResponse(response, true)
        if err != nil {
            panic(err)
        }
        fmt.Println(string(dump))
        // 全部送信完了していれば終了
        current++
        if current == len(sendMessages) {
            break
        }
    }
    conn.Close()
}
