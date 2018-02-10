package s2_2

import(
    "bufio"
    "fmt"
    "io/ioutil"
    "net"
    "net/http"
    "net/http/httputil"
    "os"
    "path/filepath"
    "strings"
)

// Unixドメインソケット版のHTTPサーバーを作る
func main() {
    path := filepath.Join(os.TempDir(), "unixdomainsocket-sample")
    os.Remove(path)
    listener, err := net.Listen("unix", path)
    if err != nil {
        panic(err)
    }
    defer listener.Close()
    fmt.Println("Server is runnning at " + path)
    for {
        conn, err := listener.Accept()
        if err != nil {
            panic(err)
        }
        go func() {
            fmt.Printf("Accept %v\n", conn.RemoteAddr())
            // リクエストを読み込む
            request, err := http.ReadRequest(bufio.NewReader(conn))
            if err != nil {
                panic(err)
            }
            dump, err := httputil.DumpRequest(request, true)
            if err != nil {
                panic(err)
            }
            fmt.Println(string(dump))
            // レスポンスを書き込む
            response := http.Response {
                StatusCode: 200,
                ProtoMajor: 1,
                ProtoMinor: 0,
                Body: ioutil.NopCloser(
                    strings.NewReader("Hello world\n")),
            }
            response.Write(conn)
            conn.Close()
        }()
    }
}
