package main

import (
	"fmt"

	"github.com/don0s/go-practice/pkg/datastructures"
)

func main() {
	fmt.Println("Hello world")

	list := datastructures.CreateLinkedList("C")
	datastructures.AddToHead(list, "B")
	datastructures.AddToHead(list, "A")

	datastructures.AddToTail(list, "D")
	datastructures.AddToTail(list, "E")
	fmt.Println(list)

}
