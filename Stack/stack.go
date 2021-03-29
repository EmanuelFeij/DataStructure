package stack
//Stack type definition
type Stack struct {
	items []int
}

//Push one item to the stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop the last item from the stack
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return -1, nil
	}
	l := len(s.items) - 1
	lastItem := s.items[l]
	s.items = s.items[:l]
	return lastItem