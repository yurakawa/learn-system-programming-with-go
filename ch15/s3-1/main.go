package main

import (
    "fmt"
)

// スライスの作成方法
func main() {
    // 既存の配列を参照するスライス
    a := [4]int{1, 2, 3, 4}
    b := a[:]
    fmt.Println("b:", &b[0], len(b), cap(b))

    // 既存の配列の一部を参照するスライス
    c := a[1:3]
    fmt.Println("c:",&c[0], len(c), cap(c))

    // 何も参照していないスライス
    //   構造体のメンバー変数などではこの状態で末初期化にしておいてメモリを遅延確保するという使い方をする
    var d []int
    fmt.Println("d:", len(d), cap(d))
    // 0 0

    // 初期の配列とリンクされているスライス
    e := []int{1, 2, 3, 4}
    fmt.Println("e:", &e[0], len(e), cap(e))

    // サイズを持ったスライスを定義
    f := make([]int, 4)
    fmt.Println("f:", &f[0], len(f), cap(f))

    g := make([]int, 4, 8)
    fmt.Println(&g[0], len(g), cap(g))
}


