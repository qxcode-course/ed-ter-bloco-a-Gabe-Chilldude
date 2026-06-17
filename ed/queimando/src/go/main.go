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

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()

	stack.Push(Pos{l:l,c:c})
	for _ = 0; !(stack.IsEmpty()); _ = 0  {
		pos := stack.Pop()
		if pos.l >= 0 && pos.l < len(grid) &&
			pos.c >= 0 && pos.c < len(grid[0]) {
				if grid[pos.l][pos.c] == '#' {

					grid[pos.l][pos.c] = 'o'
					stack.Push(Pos{l:pos.l+1,c:pos.c})
					stack.Push(Pos{l:pos.l-1,c:pos.c})
					stack.Push(Pos{l:pos.l,c:pos.c+1})
					stack.Push(Pos{l:pos.l,c:pos.c-1})
				}
			}
	}
	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
