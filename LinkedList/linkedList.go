package LinkedList

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

func (l *LinkedList) Add(data int) {
	newNode := &node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		l.capacity = 1
		return
	}
	var current *node
	for i := l.head; i != nil; i = i.next {
		current = i
	}
	current.next = newNode
	l.capacity++
	return
}

func (l *LinkedList) Remove(data int) {
	if l.head != nil {
		if l.head.data == data {
			l.head = l.head.next
			l.capacity--
		} else {
			var previous *node
			for i := l.head; i != nil; i = i.next {
				if i.data == data {
					previous.next = i.next
					l.capacity--
				}
				previous = i
			}
		}

	}

}

func (l *LinkedList) ToSlice() []int {
	slice := make([]int, 0)
	for i := l.head; i != nil; i = i.next {
		slice = append(slice, i.data)
	}
	return slice
}
