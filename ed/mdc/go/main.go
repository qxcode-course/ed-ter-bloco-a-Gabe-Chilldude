package main

import (
	"fmt"
)

func mdc(a, b int) int {
	if a == 0 || b == 0 { 
		return a+b
	}
	
	return mdc(b, a%b)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
