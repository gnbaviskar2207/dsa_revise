package linkedlists

import "fmt"

type LinkedList struct {
	Head *Node
	Size int
	Tail *Node
}

func (ll *LinkedList) InsertAtHead(val int) {
	newNode := NewDefaultNode(val)
	newNode.Next = ll.Head
	ll.Head = newNode
	ll.Tail = newNode
	ll.Size += 1
}

func (ll *LinkedList) InsertAtTail(val int) {
	curr := ll.Head
	node := NewDefaultNode(val)

	if ll.Head == nil {
		ll.InsertAtHead(val)
		return
	}

	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
	ll.Tail = node
	ll.Size += 1
}

func (ll *LinkedList) PrintLinkedList() {
	curr := ll.Head
	fmt.Print("{ ")
	for curr != nil {
		fmt.Printf("%d  =>  ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("  ||  Head  =>  %d  Tail  =>  %d } \n", ll.Head.Val, ll.Tail.Val)
}

func NewDeailtLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Size: 0,
		Tail: nil,
	}
}
