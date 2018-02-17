package main

import (
    "context"
    "fmt"
    "github.com/lestrrat/go-server-starter/listener"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

// go build server.go
// start_server --port 8080 --pid-file app.pid -- ./server

// Server::Starter対応のサーバの実装例
func main () {
    // シグナル初期化
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, syscall.SIGTERM)

    // Server::Starterから貰ったソケットを確認
    listeners, err := listener.ListenAll()
    if err != nil {
        panic(err)
    }
    if err != nil {
        panic(err)
    }

    // ウェブサーバをgoroutineで起動
    server := http.Server {
        Handler : http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, "server pid: %d %v\n", os.Getpid(), os.Environ())
        }),
    }
    go server.Serve(listeners[0])

    // SIGTERM を受け取ったら終了させる
    <-signals
    server.Shutdown(context.Background())
}
