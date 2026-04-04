package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func soma(vet []int) []int {
	_=vet
	var res = make([]int, 0)
	if len(vet) == 2 {
		var aux = make([]int, 0)
		aux = append(res, vet[0] + vet[1])
		return aux 
	}
	
	var ult = len(vet)-1
	var value = vet[ult-1] + vet[ult]
	res = soma(vet[:ult])
	res = append(res, value)

	return res
}

func processa(vet []int) {
	_ = vet
	// 1. defina o ponto de parada
	// 2. monte o vetor auxiliar com os resultados das somas
	// 3. chame recursivamente a função processa para o vetor auxiliar
	// 4. imprima o vetor original
	if len(vet) == 1 {
		fmt.Printf("[ %d ]\n", vet[0])
		return
	}
	processa(soma(vet))
	fmt.Print("[")
	for i := range len(vet) {
		fmt.Printf(" %d", vet[i])
	}
	fmt.Print(" ]\n")
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}
