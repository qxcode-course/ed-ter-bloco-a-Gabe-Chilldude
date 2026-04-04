package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	_, _ = x, div

	if x == 1 {
		return false
	}
	if x != div && x%div == 0 {
		return false
	}
	if x == div {
		return true
	}

	res := eh_primo(x, div+1)

	return res;
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
