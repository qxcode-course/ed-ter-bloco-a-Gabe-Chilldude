package main
import "fmt"

func resto(num int) {
	var remain = num % 2
	var result = num/2
	if num == 1 {
		fmt.Println(result, remain)
		return
	}

	resto(result)
	fmt.Println(result, remain)
}

func main() {
	var num int
	fmt.Scan(&num)
	resto(num)
}
