package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "net"
    "net/http"
    "net/http/httputil"
    "strings"
    "time"
    "io"
)

/**
* パイプライニングのサーバ実装
* 仕様: サーバ皮の状態を変更しないメソッド(GETやHEAD)であれば、サーバ側で並列処理を行って良い
*       リクエストの順序でレスポンスを返さなければならない
*/

// 順番に従ってconnに書き出しをする(goroutineで実行される)
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
    defer conn.Close()
    // 順番に取り出す
    for sessionResponse := range sessionResponses {
        // 選択された仕事が終わるまで待つ
        response := <-sessionResponse
        response.Write(conn)
        close(sessionResponse)
    }

}

func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
    dump, err := httputil.DumpRequest(request, true)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(dump))
    content := "Hello World\n"
    // レスポンセを書き込む
    // セッションを維持するためにKeep-Aliveでないといけない
    response := &http.Response{
        StatusCode: 200,
        ProtoMajor: 1,
        ProtoMinor: 1,
        ContentLength: int64(len(content)),
        Body: ioutil.NopCloser(strings.NewReader(content)),
    }
    // 処理が終わったらチャネルに書き込み
    // ブロックされていたwriteToConnの処理を再始動する
    resultReceiver <- response
}

// セッション1つを処理
func processSession(conn net.Conn) {
    fmt.Printf("Accept %v\n", conn.RemoteAddr())
    // セッション内のリクエストを順に処理するためのチャネル
    sessionResponses := make(chan chan *http.Response, 50)
    defer close(sessionResponses)
    // レスポンスを直列化してソケットに書き出す専用のgoroutine
    go writeToConn(sessionResponses, conn)
    reader := bufio.NewReader(conn)
    for {
        // レスポンスを受け取ってセッションのキューに入れる
        conn.SetReadDeadline(time.Now().Add(5 * time.Second))
        // リクエストを読み込む
        request, err := http.ReadRequest(reader)
        if err != nil {
            neterr, ok := err.(net.Error)
            if ok && neterr.Timeout() {
                fmt.Println("Timeout")
                break
            } else if err == io.EOF {
                break
            }
            panic(err)
        }
        sessionResponse := make(chan *http.Response)
        sessionResponses <- sessionResponse
        // 非同期でレスポンスを実行
        go handleRequest(request, sessionResponse)
    }
}

func main () {
    listener, err := net.Listen("tcp", "localhost:8888")
    if err != nil {
        panic(err)
    }
    fmt.Println("Server is running at localhost:8888")
    for {
        conn, err := listener.Accept()
        if err != nil {
            panic(err)
        }
        go processSession(conn)
    }
}
