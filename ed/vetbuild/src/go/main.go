package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
	"errors"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v Vector) Status() string{
	return fmt.Sprint("size:", v.size, " capacity:", v.capacity)
}

func (v *Vector) String() string{
	res :="" 
	for i := range v.size{
		res = res + fmt.Sprint(v.data[i])
		if i < v.size-1 {
			res += ","
			res += " "
		}
	}
	return fmt.Sprint("[",res,"]")
}


func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (v *Vector) PushBack(value int) {
	auxSize := v.size + 1
	if auxSize > v.capacity {
		v.capacity *= 2
		newData := make([]int, v.capacity)
		for i := range v.size {
			newData[i] = v.data[i]
		}
		newData[auxSize-1] = value
		v.size ++
		v.data = newData
	} else {
		v.size ++
		v.data[v.size-1] = value
	}
}

func (v *Vector) Set(index, value int) error{
	if index < v.size && index >= 0 {
		v.data[index] = value
		return nil
	}
	return errors.New("index out of range")
}

func (v *Vector) Get(index int) int {
	return v.data[index]
}

func (v *Vector) At(index int) (int, error) {
	if index < v.size && index >= 0 {
		return  v.Get(index), nil
	}
	return 0, errors.New("index out of range")
}

func (v *Vector) Clear() {
	v.size = 0
	cleared := make([]int, 0)
	v.data = cleared;
}

func (v *Vector) Reserve(newCap int) {
	aux := make([]int, newCap)
	
	for i := range(v.capacity) {
		aux[i] = v.data[i]
	}

	aux = v.data
	v.capacity = newCap
}

func (v *Vector) PopBack() (int, error) {
	if v.size <= 0 {
		return 0,  errors.New("vector is empty")
	}

	v.size --

	return v.data[v.size], nil
}

func (v *Vector) Insert(index, value int) error {
	if index >= v.capacity || index < 0 {
		return errors.New("index out of range")
	}
	var aux []int

	if v.size == v.capacity {
		v.capacity *= 2
	}
	aux = make([]int, v.capacity)

	for i := range(v.size+1) {
		if i < index {
			aux[i] = v.data[i]
		} else if i == index {
			 aux[i] = value
		} else {
			aux[i] = v.data[i-1]
		}
	}

	v.size ++
	v.data = aux

	return nil
}

func (v *Vector) IndexOf(value int) int{
	for i := range(v.size) {
		if v.data[i] == value {
			return i 
		}
	}
	return -1
}

func (v *Vector) Contains(value int) bool {
	return v.IndexOf(value) > -1
}
func (v *Vector) Erase(index int) error {
	if index >= v.capacity || index < 0 {
		return errors.New("index out of range")
	}

	for i := index; i<v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}

	v.size --
	return nil
}

func (v *Vector) FitIndex(index int) int{
	if index < 0 {
		return index + v.size
	}
	return index % v.size
}

func (v *Vector) Slice(start, end int) *Vector {
	start = v.FitIndex(start)
	end = v.FitIndex(end)

	slice := NewVector(end-start)

	for i:=start;i<end;i++ {
		slice.PushBack(v.data[i])
	}
	return slice 
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0) // declaracao do vetor
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			 for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
			 	v.PushBack(value)
			}
		case "show":
			fmt.Println(v.String())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
