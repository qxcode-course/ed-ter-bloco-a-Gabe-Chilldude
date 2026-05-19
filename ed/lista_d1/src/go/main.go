package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	value int
	next *Node
	prev *Node
}

type LList struct {
	root *Node
}

func NewLList() LList{
	var root Node
	root = Node {
		next: &root,
		prev: &root,
	}

	return LList{
		root: &root,
	}
}

func insert(node *Node, value int) {
	previous := node.prev

	newNode := Node {
		value: value,
		next: node,
		prev: previous,
	}

	node.prev = &newNode
	previous.next = &newNode
}

func remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil
	node.next = nil
}

func (node Node) Next() *Node{
	return node.next
}

func (node Node) Prev() *Node{
	return node.prev
}

func (node Node) GetValue() int {
	return node.value
}

func (node * Node) SetValue(value int) {
	node.value = value
}

func (l LList) Front() *Node{
	return l.root.next
}

func (l LList) Back() *Node{
	return l.root.prev
}

func (l LList) End() *Node{
	return l.root
}

func (l * LList) PushFront(value int) {
	insert(l.Front(), value)
}

func (l * LList) PushBack(value int) {
	insert(l.End(), value)
}

func (l * LList) PopFront() {
	if l.Front() != l.End() {
		remove(l.Front())
	}
}

func (l * LList) PopBack() {
	if l.Back() != l.End() {
		remove(l.Back())
	}
}

func (l LList) Size() int{
	size := 0
	for i := l.Front(); i != l.End(); i = i.next {
		size ++
	}

	return size
}

func (l * LList) Clear() {
	// meu Deus, muito loop
	// talvez seja melhor ter um LList.size
	// ai incrementa ou decrementa em push ou pop 
	for i:= 0; l.Size() > 0; i++ {
		l.PopFront()
	}
}

func (l * LList) String() string{
	res := "["
	for i := l.Front(); i != l.End(); i = i.next {
		res = res + fmt.Sprint(i.GetValue())
		if i != l.Back() {
			res = res + ", "
		}
	}
	res += "]"
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
