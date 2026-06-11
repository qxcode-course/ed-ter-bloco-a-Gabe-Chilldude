package main
import "fmt"
// por favor deivide sena me emsine a fazer deques pilhas e filas
// queue : inserção em uma ponta e remoção em outra
// stack : inserção em um ponta e remoção na mesma ponta
// deque : double ended queue
	// front : primeiro elemento do deque
	// size : tamanho do deque
	// cap : capacidade do deque

type Deque[T comparable] struct {
	value []T
	front * int
	size int
	capacity int
}

func newDeque[T comparable]() *Deque[T]{

	return nil
}

func main() {
    fmt.Println("Hello, World!")
}
