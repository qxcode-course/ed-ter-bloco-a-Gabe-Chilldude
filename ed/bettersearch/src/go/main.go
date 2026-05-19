package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BinarySearch(high, low, value int, slice []int) (bool, int){
	if high >= len(slice) || low >= len(slice) ||
	high < 0 || low < 0 {
		return false, ()
	}
	if high == low {
		return true, high-low
	}
	var newHigh, newLow int
	if slice[high-low] == value {
		return true, high-low
	}
	if slice[high-low]  > value {
		newHigh = high/2
	}
	if slice[high-low]  < value {
		newLow = low*2
	}

	return BinarySearch(newHigh, newLow, value, slice)
}

func BetterSearch(slice []int, value int) (bool, int) {
	_, _ = slice, value
	
	return BinarySearch(len(slice)-1, 0, value, slice)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
