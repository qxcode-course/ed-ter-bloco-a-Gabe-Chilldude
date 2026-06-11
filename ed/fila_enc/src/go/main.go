package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


type Node[T any] struct {
	Value T
	next  *Node[T]
}


type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	queque := &Queue[T]{
		head : &Node[T] {},
	} 
	return queque
}

func (q *Queue[T]) Enqueue(value T) {
	elem := &Node[T] { Value : value, }
	if q.tail == nil {
		q.head = elem
		q.tail = elem
		return
	}
	
	q.tail.next = elem
	q.tail = elem
}

func (q *Queue[T]) IsEmpty() bool {
	return q.tail == nil
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var r T
		return r, false
	}
	if q.head.next != nil {
		q.head = q.head.next
	}
	q.size -- 
	return q.head.next.Value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		return q.head.Value, false
	}

	return q.head.Value, true
}

func (q *Queue[T]) Size() int {
	return q.size
}
func (q *Queue[T]) Clear() {
	q.head.next = nil
	q.size = 0
}

func (q *Queue[T]) String() string {
	result := "["
	for n := q.head; n != nil; n = n.next {
		if n != q.head {
			result += ", "
		}
		result += fmt.Sprintf("%v", n.Value)
	}
	return result + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[int]()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "end":
			break
		case "show":
			fmt.Println(queue)
		case "push":
			for _, arg := range args[1:] {
				value, _ := strconv.Atoi(arg)
				queue.Enqueue(value)
			}
		case "pop":
			if _, ok := queue.Dequeue(); !ok {
				fmt.Println("falha: fila vazia")
			}
		case "peek":
			if value, ok := queue.Peek(); ok {
				fmt.Println(value)
			} else {
				fmt.Println("falha: fila vazia")
			}
		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
