package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func obejtosParaString(vet []int) string {
	_ = vet
	if len(vet) == 1 {
		return strconv.Itoa(vet[0])
	} else if len(vet) == 0 {
		return ""
	}
	res := strconv.Itoa(vet[0]) + ", " + obejtosParaString(vet[1:])

	
	return res
}

func tostr(vet []int) string {
	res := obejtosParaString(vet)
	return "[" + res + "]"
}

func tostrrev(vet []int) string {
	_ = vet
	res := obejtosParaString(reverse(vet))
	return "[" + res + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) []int{
	_ = vet
	if len(vet) <= 1 {
		return vet
	}
	fim  := len(vet) - 1
	vet[0], vet[fim] = vet[fim], vet[0]
	reverse(vet[1:fim])

	return vet
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	_ = vet
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	_ = vet
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
	return 0
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	_ = vet
	if len(vet) == 0 {
		return -1
	}
	var indiceMin func(v []int, i int) (int, int)
	indiceMin = func(v []int, i int) (int, int){
		if len(v) == 1 {
			return v[0], i
		}
		res, index := indiceMin(v[1:], i+1)
		if v[0] < res {
			return v[0], i
		} 

		return res, index
	}

	var  _, res = indiceMin(vet, 0)
	return res
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
