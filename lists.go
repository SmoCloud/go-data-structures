package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	number *int
	next   *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) Push(n int) {
	newNode := &Node{number: &n, next: nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = l.tail.next
	}
	l.size++
}

func (l *LinkedList) Pop() int {
	current := l.head
	val := *l.tail.number
	for current.next != l.tail {
		current = current.next
	}
	l.tail = current
	l.size--
	return val
}

func (l *LinkedList) Traverse() {
	current := l.head
	for range l.size {
		fmt.Println(*current.number)
		current = current.next
	}
}

func main() {
	list := LinkedList{head: nil, tail: nil}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for range 10 {
		list.Push(rand.Int() % 100)
	}
	fmt.Println("List size: ", list.size)
	list.Traverse()
}
