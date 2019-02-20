package main

import (
	"fmt"
	"image/color"
)

//IntList is a linked list Integers
type IntList struct {
	Value int
	Tail  *IntList
}

//Sum calcuates the sume of list
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

//ColoredPoint is
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	fmt.Println("main() start")
	fmt.Println("main() end")
}
