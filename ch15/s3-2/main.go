package main

import(
    "fmt"
)

// スライスのメモリ確保
func main(){
    // 長さ1, 確保された要素2のスライスを作成
    s := make([]int, 1, 2)
    fmt.Println(&s[0], len(s), cap(s))
    fmt.Println(s)

    // 1要素追加
    s = append(s, 100)
    fmt.Println(&s[0], len(s), cap(s))
    fmt.Println(s)

    // 更に要素を追加(新しく配列が確保され直す)
    s = append(s, 200)
    fmt.Println(&s[0], len(s), cap(s))
    fmt.Println(s)
    // スライスの先頭要素のアドレスも変わる

}
