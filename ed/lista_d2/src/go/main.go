package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node[T comparable] struct {
	Value T
	root *Node[T]
	next *Node[T]
	prev *Node[T]
}

type LList[T comparable] struct {
	root *Node[T]
}

func NewLList[T comparable]() LList[T]{
	var root Node[T]
	root = Node[T] {
		root: &root,
		next: &root,
		prev: &root,
	}

	return LList[T] {
		root: &root,
	}
}

func (l *LList[T]) Insert(node *Node[T], value T) *Node[T]{
	previous := node.prev

	newNode := Node[T] {
		Value: value,
		root: l.root,
		next: node,
		prev: previous,
	}

	node.prev = &newNode
	previous.next = &newNode

	return &newNode
}

func (l *LList[T]) Remove(node *Node[T]) *Node[T]{
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node.prev = nil
	node.next = nil

	return next
}

func (node Node[T]) Next() *Node[T]{
	if node.next != node.root {
		return node.next
	}
	return nil
}

func (node Node[T]) Prev() *Node[T]{
	if node.prev != node.root {
		return node.prev
	}
	return nil
}

func (l LList[T]) Front() *Node[T]{
	if l.root.next != l.root {
		return l.root.next
	}
	return nil
}

func (l LList[T]) Back() *Node[T]{
	if l.root.prev != l.root {
		return l.root.prev
	}
	return nil
}

func (l LList[T]) End() *Node[T]{
	return l.root
}

func (l * LList[T]) PushFront(value T) {
	l.Insert(l.Front(), value)
}

func (l * LList[T]) PushBack(value T) {
	l.Insert(l.End(), value)
}

func (l * LList[T]) PopFront() {
	if l.Front() != nil {
		l.Remove(l.Front())
	}
}

func (l * LList[T]) PopBack() {
	if l.Back() != nil {
		l.Remove(l.Back())
	}
}

func (l LList[T]) Size() int{
	size := 0
	for i := l.Front(); i != nil; i = i.Next() {
		size ++
	}

	return size
}

func (l * LList[T]) Clear() {
	for i:= l.Front(); i != nil; i = l.Front() {
		l.Remove(i)
	}
}

func (l * LList[T]) Search(value T) *Node[T]{
	for i:= l.Front(); i != nil; i= i.Next() {
		if i.Value == value {
			return i
		}
	}
	return nil
}

func (l * LList[T]) String() string{
	res := "["
	for i := l.Front(); i != l.End(); i = i.next {
		res = res + fmt.Sprint(i.Value)
		if i != l.Back() {
			res = res + ", "
		}
	}
	res += "]"
	return res
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
   ll := NewLList[int]()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
