package main
import "fmt"

func scanArray(array []int) {
	for i := range array {
		fmt.Scan(&array[i])
	}
}

func prettyArray(array []int) {
	for i:=0;i < len(array);i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Print("\n")
}

func arrayRemove(array []int, targets []int) []int{
	for i:= range array {
		for j:= range targets {
			if array[i] == targets[j] {
				array[i] = 0
			}
		}
	}
	
	newArray := make([]int, 0)
	for i:= range array {
		if array[i] != 0{
			newArray = append(newArray, array[i])
		}
	}

	return newArray
}

func main() {
	var inicio,saidas int
	
	fmt.Scan(&inicio)

	fila := make([]int, inicio)

	scanArray(fila)

	fmt.Scan(&saidas)

	desertores := make([]int, saidas)

	scanArray(desertores)

	filaFinal := arrayRemove(fila, desertores)

	prettyArray(filaFinal)

}
