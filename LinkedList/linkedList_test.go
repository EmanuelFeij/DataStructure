package LinkedList

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head     *node
	capacity int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, capacity: 0}
}

func (l *LinkedList) add(data int) {
	newNode := &node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	var current *node
	for i := l.head; i != nil; i = i.next {
		current = i
	}
	current.next = newNode
	return
}

func (l *LinkedList) print() {
	for i := l.head; i != nil; i = i.next {
		fmt.Println(i.data)
	}
}
