package main
import (
	"fmt"
)

// estrategia : valores temporarios
// na ida construir um array de indices ocupados ( que não possuem .) e na volta testar todos os indices com 
// diferença >=3 do indice atual e mudar desta maneira
// o complicado é pegar o valor ... DESCOBRI

// TODO fazer com que ele resolva casos com muitos espaços vazios
func solve(nlist []rune, occ []int, pos int, maxi rune) []int{
	maxiInt := int(maxi - '0')

	if pos >= len(nlist) || pos < 0 { // checa se posicao esta dentro de nlist
		return occ
	}

	if nlist[pos] != '.' { // checa se posicao em nlist não é '.' ou seja ocupado por um digito
		occ = append(occ, pos) // insere obj da pos atual em ocupados
	}

	occ = solve(nlist, occ, pos+1, maxi)

	curr := maxi // valor a ser usado em um espaco vazio
	fmt.Printf("%d %d\n", occ, pos)

	if nlist[pos] == '.' {
		fmt.Println("comeca", pos)
		for { // roda até achar resposta (se não existir, o algoritmo ta errado (?))
			unq := true // unique
			for _, i := range occ {
				if i < pos-maxiInt || i > pos+maxiInt { // checa se i esta em L posicoes de distancia de pos 
					fmt.Printf("0. %d %c, pos: %d\n", pos, nlist[i], i)
					continue
				}
				fmt.Printf("1. %c %c, pos: %d\n", curr, nlist[i], i)
				if nlist[i] == curr { // se valor a ser inserido for igual a um valor em distancia L
					fmt.Println(curr, " moio :(")
					unq = false // nao e unico
					break // proximo loop
				} else {
					unq = true
				}
			}

			if unq {
				break
			}
			 
			curr = curr - 1  // incrementa para proximo loop
			if curr < '0'  {
				curr = '9'
			} else if curr > '9' {
				curr = '0'
			}
		}
		nlist[pos] = curr // substitui valor de '.' pelo valor de curr na lista
		occ = append(occ, pos) // adiciona valor da pos em nList para ocupados
	}
	
	return occ
}

func main() {
	var linha1 string
	var linha2 int
	
	fmt.Scanf("%s", &linha1)
	fmt.Scanf("%d", &linha2)

	nlist := []rune(linha1)
	occ := make([]int, 0)

	solve(nlist, occ, 0, rune(linha2)+'0')
	
	fmt.Printf("%c\n", nlist)
}
