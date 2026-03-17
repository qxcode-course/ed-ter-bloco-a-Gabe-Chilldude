package main
import "fmt"
func main() {
	var n int
	var casais int
	fmt.Scan(&n)
	var descasados = make([]int, n)

	for i:=0;i<n;i++{
		fmt.Scan(&descasados[i])
	}

	for i:=0;i<n;i++{
		if descasados[i] == 0 {
			continue
		}
		for j:=0;j<n;j++{
			if descasados[j] == -(descasados[i]) {
				descasados[j] = 0
				descasados[i] = 0
				casais ++
			}
		}
	}

	fmt.Println(casais)
}
