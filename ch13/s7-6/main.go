package main

import (
    "sync"
    "fmt"
)

// sync.Map 複数のgoroutineからアクセスされても壊れない
func main () {
    smap := sync.Map{}
    smap.Store("hello", "world")
    smap.Store(1, 2)
    smap.Store(3, 2)

    // 削除
    smap.Delete(1)

    // 取り出し方法
    value, ok := smap.Load("hello")
    fmt.Printf("key=%v value=%v exists?=%v\n", "hellow", value, ok)


    // キーが登録されていれば過去のデータを、登録されていなければ新しい値を登録する
    smap.LoadOrStore(3,  3) // すでに保存されているので無視
    smap.LoadOrStore(4, 2) // 保存されていないので挿入

    // ループで取り出し
    smap.Range(func(key, value interface{}) bool {
        fmt.Printf("%v: %v\n", key, value)
        return true
    })
}

