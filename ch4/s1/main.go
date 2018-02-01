package main

import (
	"fmt"
	"time"
)

// 新しく作られる goroutineが呼ぶ関数
func sub() {
	fmt.Println("sub() is runnning")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func main() {
	fmt.Println("start sub ()")
	// goroutine を作って関数を実行
	go sub()
	time.Sleep(2 * time.Second)
}
