package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// チャネルを作ってsignal.Notify()に私、OSから受け取ったシグナルをこのチャネルに送信する
func main() {
	// サイズが1より大きいチャネルを作成
	signals := make(chan os.Signal, 1)
	// SIGINT (Ctrl+C) を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルが来るまで待つ
	fmt.Println("Waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")

}
