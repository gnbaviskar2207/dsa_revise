package stackquque

type StackLL struct {
	Top  *Node
	Size int
}

func (s *StackLL) push(ele int) {
	node := NewDefaultNode(ele)
	node.Next = s.Top
	s.Top = node
	s.Size += 1
}

func (s *StackLL) pop() (int, bool) {
	if s.Top == nil {
		return 0, false
	}
	val := s.Top.Val
	temp := s.Top.Next
	s.Top.Next = nil
	s.Top = temp
	s.Size -= 1
	return val, true
}

func (s *StackLL) top() (int, bool) {
	if s.Top == nil {
		return 0, false
	}
	return s.Top.Val, true
}

func NewDefaultStackLL() *StackLL {
	return &StackLL{
		Size: 0,
		Top:  nil,
	}
}
