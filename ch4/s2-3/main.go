package main

import (
	"fmt"
	"math"
)

// エラトステネスの篩で素数を算出
func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			// 3以上の奇数
			for j := 3; j < l + 1; j +=2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primeNumber()
	// ここがポイント
	for n := range pn {
		fmt.Println(n)
	}
}
