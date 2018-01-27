package main

import (
	"os"
	"fmt"
	"time"
	"encoding/json"
)

func main() {
	// %vはなんでも表示できるフォーマット指定子
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())

	// JSONを整形して出力する
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("","      ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello": "world",
	})
}
