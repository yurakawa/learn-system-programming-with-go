package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

type Id int64

func (id *Id) Inc() {
    atomic.AddInt64((*int64)(id), 1)
}

// s7-1をベースに sync/atomicやsync.WaitGroupを組み込む
func main() {
    var id Id
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            id.Inc()
            fmt.Printf("id: %d\n",id)
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println("Count =", id, "GOMAXPROCS =", runtime.GOMAXPROCS(0))
}