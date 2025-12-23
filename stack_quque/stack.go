package stackquque

type Stack struct {
	index int
	size  int
	arr   []int
}

/*
* TC: O(1)
 */
func (s *Stack) push(ele int) {
	if s.index >= s.size-1 {
		panic("stack overflow")
	}
	s.index += 1
	s.arr[s.index] = ele
}

/*
* TC: O(1)
 */
func (s *Stack) pop() int {
	if s.index < 0 {
		panic("stack underflow")
	}
	ele := s.arr[s.index]
	s.index -= 1
	return ele
}

/*
* TC: O(1)
 */
func (s *Stack) top() int {
	if s.index < 0 {
		panic("stack underflow")
	}
	if s.index > s.size-1 {
		panic("stack overflow")
	}
	return s.arr[s.index]
}

/*
* TC: O(1)
 */
func (s *Stack) isEmpty() bool {
	return s.index == -1
}

/*
* TC: O(1)
 */
func (s *Stack) isFull() bool {
	return s.index == s.size-1
}

func NewDefaultStack(size int) *Stack {
	return &Stack{
		index: -1,
		size:  size,
		arr:   make([]int, size),
	}
}
