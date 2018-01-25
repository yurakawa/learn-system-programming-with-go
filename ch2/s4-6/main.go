package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	// %vはなんでも表示できるフォーマット指定子
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
