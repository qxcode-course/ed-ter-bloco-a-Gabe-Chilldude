package main
import "fmt"
func main() {
	var n, m int

	fmt.Scan(&n)
	filaInicio := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&filaInicio[i])
	}

	fmt.Scan(&m)
	desertores := make([]int, m)
	for i:=0; i<m; i++ {
		fmt.Scan(&desertores[i])
	}

	filaFim := make(map[int]int)

	// remover desertores do array original

	filaContador := 0
	for i:=0;i<n;i++ {
		saiuDaFila := false
		for j:=0;j<m;j++ {
			if filaInicio[i] == desertores[j] {
				saiuDaFila = true
				break
			}
		}
		if !saiuDaFila {
			filaFim[filaContador] = filaInicio[i]
			filaContador ++ 
		}
	}

	for i:=0; i < len(filaFim); i++ {
		fmt.Print(filaFim[i])
		fmt.Print(" ")
	}
	fmt.Print("\n")

}
