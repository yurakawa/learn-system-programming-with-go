package main

import (
	"fmt"
	"strconv"
	"time"
)

// select
func main() {
	fmt.Println("開始")

	ch1 := make(chan int)
	ch2 := make(chan string)
	exit := make(chan struct{}) // 終了通知用のチャネル

	// チャネルを送信するgoroutine
	go func(chint chan<- int, chstr chan<- string, end chan<- struct{}) {
		for i := 0; i < 10; i++ {
			// 偶数階はint型チャネル、奇数回はstring型チャネルを送信する
			if i%2 == 0 {
				fmt.Println("ch1へ送信")
				chint <- i
			} else {
				fmt.Println("ch2へ送信")
				chstr <- "test" + strconv.Itoa(i)
			}
		}

		time.Sleep(1 * time.Second)
		close(end) //クローズして通知
	}(ch1, ch2, exit)

	// 受信用の無限ループ
	for {
		select {
		case val := <-ch1:
			fmt.Println("ch1から受信:", val)
		case str := <-ch2:
			fmt.Println("ch2から受信:", str)
		case <-exit:
			fmt.Println("終了")
			return
		}
	}
}
