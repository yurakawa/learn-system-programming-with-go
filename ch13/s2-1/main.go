package main

import (
    "fmt"
    "time"
)

func sub1 (c int) {
    fmt.Println("share by arguments:", c*c)
}

//goroutineと情報共有
func main(){
    // 引数渡し
    go sub1(10)

    // クロージャのキャプチャ渡し
    c := 20
    go func() {
        fmt.Println("share by capture", c*c)
    }()

    // sleepがないと親スレッドが終わってしまいgoroutineが実行されない場合がある
    time.Sleep(time.Second)
}


