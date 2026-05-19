package main
import "fmt"
type Node struct {
	value int 
	next *Node
	previous *Node
}

type DList struct {
	head *Node
}

func NewDlist() DList {
	var list DList
	var head Node

	list.head = &head
	list.head.next = list.head
	list.head.previous = list.head

	return list
}

func insert(node *Node, value int)  {
	previous := node.previous

	newNode  := Node{
		value: value,
		next: node,
		previous: previous,
	}

	previous.next = &newNode
	node.previous = &newNode
}

func PushBack(list *DList, value int) {
	insert(list.head.next, value)
}
func PushFront(list *DList, value int) {
	insert(list.head, value)
}

func main() {
	fmt.Println("")
}
