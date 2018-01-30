package main

import (
	"fmt"
	"strings"
)

var source = "123 1.234 1.0e4 test"

func main(){
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
