package LinkedList

import (
	"testing"
)

func TestLinkedListCreationAndAdd(t *testing.T) {
	llist := NewLinkedList()
	llist.Add(1)
	llist.Add(2)
	llist.Add(3)
	expected := []int{1, 2, 3}
	actual := llist.ToSlice()
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestLinkedListRemove(t *testing.T) {

	llist := NewLinkedList()
	llist.Add(1)
	llist.Add(2)
	llist.Add(3)
	llist.Add(5)
	llist.Add(6)
	llist.Add(7)
	llist.Remove(7)
	llist.Remove(1)
	llist.Remove(3)
	expected := []int{2, 5, 6}
	actual := llist.ToSlice()

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}
