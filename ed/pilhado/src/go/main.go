package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

func inside[T comparable](grid[][]T, pos Pos)  bool{
	return pos.l >=0 && pos.l < len(grid) && pos.c >= 0 && pos.c < len(grid[0])
}

func match[T comparable](grid[][]T, pos Pos, value T) bool{
	return inside(grid, pos) && grid[pos.l][pos.c] == value
}

func getNeigh(pos Pos) []Pos{
	return []Pos {{l:pos.l,c:pos.c+1}, {l:pos.l+1,c:pos.c}, {l:pos.l,c:pos.c-1}, {l:pos.l-1,c:pos.c}}
}

func Alter[T rune](grid [][]T, pos Pos, value T) [][]T{
	newGrid := grid
	newGrid[pos.l][pos.c] = value

	return newGrid
}

func search(grid [][]rune, pos0, posF Pos) {
	caminho := NewStack[Pos]() // stack que representa o caminho da solulção
	beco := NewStack[Pos]() // stack que representa caminhos sem saída

	caminho.Push(pos0)

	for !caminho.IsEmpty() {
		topo := caminho.Top()
		// aqui deve ser usado Stack.Top() ao invés de Stack.Pop()

		grid := Alter(grid, topo, '.')

		if topo == posF {
			break
		}

		valid := make([]Pos, 0)

		for _, i := range getNeigh(topo) {
			if !match(grid, i, '#')  && !match(grid, i, '.'){
				valid = append(valid, i)
			}
		}

		if len(valid) > 0 {
			for _, i := range valid {
				caminho.Push(i)
			}
		} else {
			beco.Push(topo)
			caminho.Pop()
		}
	}

	for !beco.IsEmpty() {
		grid = Alter(grid, beco.Pop(), ' ')
	}


}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
