package objects

import (
	"fmt"
)

type Node struct {
	Number *int
	Next   *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) Push(n int) {
	newNode := &Node{Number: &n, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = l.Tail.Next
	}
	l.Size++
}

func (l *LinkedList) Pop() int {
	current := l.Head
	val := *l.Tail.Number
	for current.Next != l.Tail {
		current = current.Next
	}
	l.Tail = current
	l.Size--
	return val
}

func (l *LinkedList) Traverse() {
	current := l.Head
	for range l.Size {
		fmt.Println(*current.Number)
		current = current.Next
	}
}