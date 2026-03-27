package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	onlyMen := make([]int, 0)
	n := len(vet)

	for i:=0; i<n; i++ {
		if vet[i] > 0 {
			onlyMen = append(onlyMen, vet[i])
		}
	}


	return onlyMen
}

func getCalmWomen(vet []int) []int {
	calmWomen := make([]int, 0)
	n := len(vet)

	for i:=0; i<n; i++ {
		if vet[i] < 0 && vet[i] > -10 {
			calmWomen = append(calmWomen, vet[i])
		}
	}

	return calmWomen
}

func sortVet(vet []int) []int { // ordenacao comum
	n := len(vet)

	for i:=0; i<len(vet); i++ {
		for j:=1; j<n; j++ {
			if vet[j] < vet[j-1] {
				temp := vet[j]
				vet[j] = vet[j-1]
				vet[j-1] = temp
			}
		}
		n--
	}

	return vet
}

func sortStress(vet []int) []int { // ordenacao por valor absoluto
	n := len(vet)

	for i:=0; i<len(vet); i++ {
		for j:=1; j<n; j++ {
			if math.Abs(float64(vet[j])) < math.Abs(float64(vet[j-1])) {
				temp := vet[j]
				vet[j] = vet[j-1]
				vet[j-1] = temp
			}
		}
		n--
	}
	return vet
}

func reverse(vet []int) []int {
	n := len(vet)
	var res []int = make([]int, 0)
	j := n-1
	for i:=0; i<n ; i++ {
		res = append(res, vet[j])
		j --
	}
	return res
}

func unique(vet []int) []int {
	n := 0
	res := make([]int, n)

	for i:=0; i<len(vet);i++ {
		unique := true
		if i == 0 {
			res = append(res, vet[i])
			n ++
		}

		for j:=0; j<n; j++ {
			if vet[i] == res[j] {
				unique = false
				break
			}
		}

		if unique {
			res = append(res, vet[i])
			n++
		}
	}
	return res
}

func repeated(vet []int) []int {
	_ = vet
	res := make([]int, 0)

	for i:=0; i<len(vet);i++ {
		if vet[i] == 0 {
			continue
		}
		hasDupe := false
		for j:=0; j<len(vet);j++ {
			if vet[j] == vet[i] && j!=i {
				hasDupe = true
				break
			}
		}

		if hasDupe {
			res = append(res, vet[i])
		}

		vet[i] = 0
	}

	sortVet(res)

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

