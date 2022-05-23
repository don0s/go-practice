package main

//"github.com/don0s/go-practice/pkg/datastructures"

import (
	"fmt"

	"github.com/don0s/go-practice/pkg/webserver"
)

func main() {

	/*
		list := datastructures.CreateLinkedList("C")
		datastructures.AddToHead(list, "B")
		datastructures.AddToHead(list, "A")

		datastructures.AddToTail(list, "D")
		datastructures.AddToTail(list, "E")
		fmt.Println(list)
	*/
	fmt.Println("Starting server...")
	webserver.StartServerMutex()
	fmt.Println("Stopping server.")
}
