package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%%s('10'): %s, %%d(10): %d, %%f(10.0): %f\n",  "10", 10, 10.0)
}
