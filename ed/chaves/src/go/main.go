package main

import (
	"fmt"
)

type Partida struct {
	plyr rune
	pnts int
}

func vencedor(participantes Queue[rune], pontos Queue[int]) rune{
	_,_ = participantes, pontos 

	for _ = 0 ; participantes.items.Len() > 1 ; _ = 0 {
		al := &Partida {
			plyr : participantes.Dequeue(),
			pnts : pontos.Dequeue(),
		}
		br := &Partida {
			plyr : participantes.Dequeue(),
			pnts : pontos.Dequeue(),
		}

		if al.pnts > br.pnts {
			participantes.Enqueue(al.plyr)
		} else {
			participantes.Enqueue(br.plyr)
		}
	}

	return participantes.Dequeue()
}

func main() {
	// fila de participantes
	participantes := NewQueue[rune]()
	var a rune = 65
	for range(16) {
		participantes.Enqueue(a)
		a ++
	}

	// recebe pontuações
	pontos := NewQueue[int]() 

	for range(15) {
		var a int
		fmt.Scan(&a)
		pontos.Enqueue(a)
		fmt.Scan(&a)
		pontos.Enqueue(a)
	}
	w := vencedor(*participantes,*pontos)
	fmt.Printf("%c\n", w)
}
