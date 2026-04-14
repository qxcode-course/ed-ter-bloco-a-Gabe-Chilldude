package main
import "fmt"

func repetidas(figs []int) []int{
	repetidas := make([]int, 0)
	for i:=1; i<len(figs); i++{
		if figs[i] == figs[i-1] {
			repetidas = append(repetidas, figs[i])
			figs[i-1] = 0
		}
	}

	return repetidas
}

func faltando(album, figs []int) []int{
	faltando := make([]int, 0)
	for i := range album {
		falta := true
		for j := range figs {
			if figs[j] == album[i] {
				falta = false
			}
		}

		if falta {
			faltando = append(faltando, album[i])
		}
	}

	return faltando
}

func prettyArray(array []int) {
	n := len(array)-1
	for i:=0;i < len(array)-1;i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Print(array[n],"\n")
}

func main() {
	var total, qtd int

	fmt.Scan(&total, &qtd)

	album := make([]int, total)
	for i := range album {
		album[i] = i+1
	}

	figs := make([]int, qtd)
	for i := range figs {
		fmt.Scan(&figs[i])
	}

	repetidas := repetidas(figs)
	faltando := faltando(album, figs)

	if len(repetidas) > 0{
		prettyArray(repetidas)
	} else {
		fmt.Println("N")
	}

	if len(faltando) > 0{
		prettyArray(faltando)
	} else {
		fmt.Println("N")
	}



}
