package Queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	v, _ := q.Pop()
	if v != 1 {
		t.Errorf("Expected 1, got %v", v)
	}
	isEmpty := q.IsEmpty()
	if isEmpty != false {
		t.Errorf("Expected false, got %v", isEmpty)
		return
	}
	expected := []int{2, 3}
	actual := q.data
	if len(expected) != len(actual) {
		t.Errorf("Expected %v , got %v", expected, actual)
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v , got %v", expected, actual)
			return
		}
	}
	return

}
