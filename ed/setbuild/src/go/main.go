package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

type Set struct {
	data []int
	size int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set {
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (v Set) String() string{
	var res string
	res = "["
	for i := range(v.size) {
		res = res + fmt.Sprint(v.data[i])
		if i < v.size-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

func (v * Set) reserve(newCapacity int) {
	v.capacity = newCapacity
	aux := make([]int, newCapacity)

	for i := range(v.size) {
		aux[i] = v.data[i]
	}
	
	v.data = aux
}

func (v * Set) insert(value, index int) {
	if v.capacity == 0 {
		v.reserve(1)
		v.size ++
	} else if v.size+1 > v.capacity {
		v.reserve(v.capacity*2)
		v.size ++
	} else {
		v.size ++
	}

	for i:=v.size-1;i>index;i-- {
		v.data[i] = v.data[i-1]
	}
	 
	v.data[index] = value
}

func (v * Set) Insert(value int) {
	var index int
	for i := range(v.size) {
		if value < v.data[i] {
			index = i
			break
		}
		if value == v.data[i] {
			return
		}
		index = i+1
	}

	v.insert(value, index)
}

func (v Set) inside(index int) bool {
	return index < v.size && index >= 0
}

func (v Set) binarysearch(value, index int) int{
	if !v.inside(index){
		return -1
	}
	if v.data[index] == value {
		return index
	}
	var newIndex int
	if v.data[index] > value {
		newIndex = (v.size-1)/4
	} else {
		newIndex = (v.size-1 - index)
		newIndex += newIndex/2
	}
	if index <= 1 {
		newIndex = 0
	}
	if newIndex == index {
		return -1
	}

	return v.binarysearch(value, newIndex)
}

func (v Set) Contains(value int) bool {
	return v.binarysearch(value, (v.size-1)/2) >= 0
}

func (v * Set) erase(index int) error{
	if !v.inside(index) {
		return errors.New("value not found")
	}
	for i:=index;i<v.size-1;i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	return nil
}

func (v * Set) Erase(value int) error {
	return v.erase(v.binarysearch(value, (v.size-1)/2))
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	_ = v
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
