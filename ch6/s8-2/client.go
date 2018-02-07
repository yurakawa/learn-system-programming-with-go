package main

import (
    "bufio"
    "fmt"
    "net"
    "net/http"
    "net/http/httputil"
    "io"
    "strconv"
)

// チャンク形式のクライアントの実装
func main() {
    conn, err := net.Dial("tcp", "localhost:8888")
    if err != nil {
        panic(err)
    }
    request, err := http.NewRequest("GET", "http://localhost:8888", nil)
    if err != nil {
        panic(err)
    }
    err = request.Write(conn)
    if err != nil {
        panic(err)
    }
    reader := bufio.NewReader(conn)
    response, err := http.ReadResponse(reader, request)
    if err != nil {
        panic(err)
    }
    dump, err := httputil.DumpResponse(response, false)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(dump))

    if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
        panic("wrong transfer encoding")
    }
    // チャンクごとに読み込む
    for {
        // サイズを取得
        sizeStr, err := reader.ReadBytes('\n')
        if err == io.EOF {
            break
        }
        // 16進数のサイズをパース。サイズがゼロならクローズ
        size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]),16,64)
        if size == 0 {
            break
        }
        if err != nil {
            panic(err)
        }
        // サイズ数分バッファを確保して読み込み
        line := make([]byte, int(size))
        reader.Read(line)
        reader.Discard(2)
        fmt.Printf(" %d bytes: %s\n", size, string(line))
    }
}
