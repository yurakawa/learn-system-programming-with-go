package main

import (
	"time"
	"fmt"
)

// time.After(duration)を使って、決まった時間を図るタイマーを作る
func main() {
	wait := make(chan string)

	select {
	case <- wait:
	case <- time.After(5 * time.Second):
		fmt.Println(5, "seconds elapsed")
	}
}

