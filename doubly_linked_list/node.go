package doublylinkedlist

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func NewDefaultNode(val int) *Node {
	return &Node{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}
