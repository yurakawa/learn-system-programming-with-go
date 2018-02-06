package main

import (
    "bufio"
    "fmt"
    "net"
    "net/http"
    "net/http/httputil"
    "strings"
    "compress/gzip"
    "os"
    "io"
)

// gzip 圧縮に対応したクライアント
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
        // このクライアントがgzipを処理できるという表明
        request.Header.Set("Accept-Encoding", "gzip")
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
        // 結果を表示(DumpResponseは圧縮された内容を理解しないためbodyは無視する)
        dump, err := httputil.DumpResponse(response, false)
        if err != nil {
            panic(err)
        }
        fmt.Println(string(dump))
        defer response.Body.Close()

        if response.Header.Get("Content-Encoding") == "gzip" {
            reader, err := gzip.NewReader(response.Body)
            if err != nil {
                panic(err)
            }
            reader.Close()
        } else {
            io.Copy(os.Stdout, response.Body)
        }
        // 全部送信完了していれば終了
        current++
        if current == len(sendMessages) {
            break
        }
    }
    conn.Close()
}
