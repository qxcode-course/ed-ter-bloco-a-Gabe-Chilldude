package main

import (
	"bufio"
	"fmt"
	"os"
)

// roubado descaradamente do exercício de ilhas

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

// STRAGIE (francês, não sei como escreve estratégia em francês
// STRAGY (favor ler com sotaque britânico, não importa a região) : procurar começando pelas extremidades por O, e detectar ilhas usando uma função auxiliar

func search(board[][]byte, pos Pos, visited, bisited map[Pos]bool) {
	if !inside(board, pos) {
		return
	}
	if visited[pos] {
		return
	}
	visited[pos] = true
	// fmt.Println("checado pelo search: ", pos)

	if match(board, pos, 'O') {
		capture(board, pos, bisited, true)
	}
	
	for _, neigh := range getNeigh(pos) {
		search(board, neigh, visited, bisited)
	}
}

func capture(board[][]byte, pos Pos, visited map[Pos]bool, capble bool) bool{
	// fmt.Println("x : ", pos)
	if !inside(board, pos) {
		return false
	}
	if board[pos.l][pos.c] == 'X' {
		// fmt.Println("X at ", pos)
		return true
	}
	valor, existe := visited[pos]
	if existe {
		return valor
	}

	visited[pos] = capble 
	for _, neigh := range getNeigh(pos) {
		if !inside(board, neigh) {
			visited[pos] = false
			break
		}
	}

	for _, neigh := range getNeigh(pos) {
		// fmt.Println("2. ", neigh, " ", visited[pos])
		visited[pos] = capture(board, neigh, visited, visited[pos]) && visited[pos]
	}


	if visited[pos] {
		board[pos.l][pos.c] = 'X'
	}

	return visited[pos] 
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	visited := make(map[Pos]bool, 0)
	bisited := make(map[Pos]bool, 0)

	search(board, Pos {l:0,c:0}, visited, bisited)
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	// fmt.Println("-----")
	for _, row := range board {
		fmt.Println(string(row))
	}
}
