package main
import "fmt"

func quantosCasais(animais []int) int {
	res := 0
	array := animais
	for i := range array {
		if array[i] == 0 {
			continue
		}
		for j:= range array {
			if array[j] == -(array[i]) {
				res++
				array[i] = 0
				array[j] = 0
				break
			}
		}
	}

	return res
}

func main() {
	var n int

	fmt.Scan(&n)

	animais := make([]int, n)

	for i:= range animais {
		fmt.Scan(&animais[i])
	}

	fmt.Println(quantosCasais(animais))

}
