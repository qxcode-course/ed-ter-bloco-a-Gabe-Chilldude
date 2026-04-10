package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func getNeig(p Pos) []Pos {
	return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
}

func inside(grid [][]rune, p Pos) bool {
	return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]))
}

func match(grid [][]rune, p Pos, value rune) bool {
	return inside(grid, p) && grid[p.l][p.c] == value
}

// Função recursiva que tenta encontrar o caminho do início ao fim
func way(grid [][]rune, startPos, endPos Pos) bool {
	_, _, _ = grid, startPos, endPos
	if !inside(grid, startPos) || grid[startPos.l][startPos.c] == '#' || grid[startPos.l][startPos.c] == '/'{
		return false
	}
	if startPos == endPos {
		grid[startPos.l][startPos.c] = '.'
		return true
	}
	res := false
	var nextPos Pos

	grid[startPos.l][startPos.c] = '/'

	nextPos.l = startPos.l + 1
	nextPos.c = startPos.c
	res = way(grid, nextPos, endPos) || res

	nextPos.l = startPos.l - 1
	nextPos.c = startPos.c
	res = way(grid, nextPos, endPos) || res

	nextPos.l = startPos.l
	nextPos.c = startPos.c + 1
	res = way(grid, nextPos, endPos) || res

	nextPos.l = startPos.l
	nextPos.c = startPos.c - 1
	res = way(grid, nextPos, endPos) || res

	if res {
		grid[startPos.l][startPos.c] = '.'
	}

	return res
}

func search(grid[][] rune, startPos, endPos Pos) bool{
	res := way(grid, startPos, endPos)

	for i:=0;i<len(grid);i++ {
		for j:=0;j<len(grid[i]);j++ {
			if grid[i][j] == '/' {
				grid[i][j] = ' '
			}
		}
	}

	return res
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
