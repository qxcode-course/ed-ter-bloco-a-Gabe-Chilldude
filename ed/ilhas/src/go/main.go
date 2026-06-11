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

func colombo(grid[][]byte, pos Pos, visited map[Pos]bool) bool{ // encontra e mapeia ilhas
	if !match(grid, pos, 49) {
		return false
	}

	if visited[pos] {
		return false
	}
	visited[pos] = true

	grid[pos.l][pos.c] = 50

	for _, value := range getNeigh(pos) {
		colombo(grid, value, visited)
	}

	return true
}
// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	//
	visited := make(map[Pos]bool, 0)
	isles := 0
	
	for index, value := range grid {
		for sindex := range value {
			if colombo(grid, Pos{l : index, c : sindex}, visited) {
				isles ++
			}
		}
	}
	_ = grid
	return isles
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
