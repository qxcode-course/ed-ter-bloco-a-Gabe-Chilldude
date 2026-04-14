package main
import "fmt"

func proxCircular(roda []int, atual int) int{
	if atual+1 > len(roda)-1{
		return 0
	}
	return atual + 1

}

func printRoda(roda[]int, assasinoAtual int) {
	fmt.Print("[ ")
	for i := range roda {
		if roda[i] == 0 {
			continue
		}
		fmt.Print(roda[i])
		if roda[i] == roda[assasinoAtual] {
			fmt.Print("> ")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("]\n")
}

func main() {
	var qtd, assasino int
	_,_ = qtd, assasino

	fmt.Scan(&qtd)
	fmt.Scan(&assasino)

	roda := make([]int, qtd)

	for i:=0;i<qtd;i++ { // popula roda
		roda[i] = i+1
	}

	assasino-- // transforma assasino em indice                        

	mortes := 0
	printRoda(roda,assasino)

	for i:=0; i<len(roda); i = proxCircular(roda, i) {
		if mortes >= len(roda) - 1 {
			break
		}
		// matar
		for j:=proxCircular(roda,assasino);j!=assasino;j=proxCircular(roda,j) {
			if roda[j] != 0 {
				roda[j] = 0
				break
			}
		}
		// proximo assasino
		for j:=proxCircular(roda,assasino);j!=assasino;j=proxCircular(roda,j) {
			if roda[j] != 0 {
				assasino = j
				break
			}
		}
		mortes ++
		printRoda(roda, assasino)
	}
}
