package Queue

import "errors"

type Queue struct {
	data     []int
	capacity int
}

func NewQueue() *Queue {
	return &Queue{
		data:     make([]int, 0),
		capacity: 0,
	}
}

func (q *Queue) Push(n int) {
	q.data = append(q.data, n)
	q.capacity++
}

func (q *Queue) IsEmpty() bool {
	return q.capacity == 0
}

func (q *Queue) Pop() (int, error) {
	if !q.IsEmpty() {
		value := q.data[0]
		q.data = q.data[1:]
		q.capacity--
		return value, nil
	}
	return -1, errors.New("empty queue")
}

func (q *Queue) Peek() int {
	if !q.IsEmpty() {
		return q.data[0]
	}
	return -1
}
