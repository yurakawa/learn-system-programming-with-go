// Goはクラスがないが、構造体がメソッドをもてる
// コンパイラが構造体のインスタンスのポインタをインタフェース型の変数に代入するコードを見つけると構造体とインスタンスの互換性をチェックするためimplementsなどのキーワードは不要

package main

import (
	"fmt"
)

// インタフェースを定義
type Talker interface {
	Talk()
}

// 構造体を宣言
type Greeter struct {
	name string
}

// 構造体はTalkerインタフェースで定義されているメソッドを持っている
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}
func main() {
	// インタフェースの型を持つ変数を宣言
	var talker Talker
	// インタフェースを満たす構造体のポインタは代入できる
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
