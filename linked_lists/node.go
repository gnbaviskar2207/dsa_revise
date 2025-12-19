package linkedlists

type Node struct {
	Val  int
	Next *Node
}

func NewDefaultNode(val int) *Node {
	return &Node{
		Val:  val,
		Next: nil,
	}
}

func NewDefaultNodeWithNext(val int, next *Node) *Node {
	return &Node{
		Val:  val,
		Next: next,
	}
}
