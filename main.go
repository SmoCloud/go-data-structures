package main

import (
	"fmt"
	"math/rand"
	"time"
	objects "github.com/SmoCloud/go-data-structures/data_structures"
)

func main() {
	list := objects.LinkedList{Head: nil, Tail: nil}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for range 10 {
		list.Push(rand.Int() % 100)
	}
	fmt.Println("List size: ", list.Size)
	list.Traverse()
}
