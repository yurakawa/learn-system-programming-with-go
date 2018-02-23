package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)
// 依存関係のあるタスクを表現する : Future/Promise

type StringFuture struct {
    receiver chan string
    cache string
}

func NewStringFuture() (*StringFuture, chan string) {
    f := &StringFuture {
        receiver: make(chan string),
    }
    return f, f.receiver
}

func (f *StringFuture) Get() string {
    r, ok := <-f.receiver
    if ok {
        close(f.receiver)
        f.cache = r
    }
    return f.cache
}

func (f *StringFuture) Close() {
    close(f.receiver)
}

func readFile(path string) *StringFuture {
    // ファイルを読み込み、その結果を返すFutureを返す
    promise, future := NewStringFuture()
    go func() {
        content, err := ioutil.ReadFile(path)
        if err != nil {
            fmt.Printf("read error %s\n", err.Error())
            promise.Close()
        } else {
            // 約束を果たした
            future <- string(content)
        }
    }()
    return promise
}

func printFunc(futureSource *StringFuture) chan []string {
    // 文字列中の関数一覧を返す Futureを返す
    promise := make(chan []string)
    go func() {
        var result []string
        // futureが解決するまで待って実行
        for _, line := range strings.Split(futureSource.Get(), "\n") {
            if strings.HasPrefix(line, "func ") {
                result = append(result, line)
            }
        }
        // 約束を果たした
        promise <- result
    }()
    return promise
}

func countLines(futureSource *StringFuture) chan int {
    promise := make(chan int)
    go func() {
        promise <- len(strings.Split(futureSource.Get(), "\n"))
    }()
    return promise
}

func main() {
    futureSource := readFile("future_promise.go")
    futureFuncs := printFunc(futureSource)
    fmt.Println(strings.Join(<-futureFuncs, "\n"))
    fmt.Println(<-countLines(futureSource))
}
