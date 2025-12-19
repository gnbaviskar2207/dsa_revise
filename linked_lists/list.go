package linkedlists

import "fmt"

type LinkedList struct {
	Head *Node
	Size int
	Tail *Node
}

/*
Insert at head
Time O(1)
*/
func (ll *LinkedList) InsertAtHead(val int) {
	node := NewDefaultNode(val)
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		node.Next = ll.Head
		ll.Head = node
	}
	ll.Size += 1
}

/*
Insert at tail
Time O(n)
*/

func (ll *LinkedList) InsertAtTail(val int) {
	curr := ll.Head
	if ll.Head == nil {
		ll.InsertAtHead(val)
		return
	}
	for curr.Next != nil {
		curr = curr.Next
	}

	node := NewDefaultNode(val)
	curr.Next = node
	ll.Tail = node
	ll.Size += 1
}

/*
Insert at tail optimized
Time O(1)
*/
func (ll *LinkedList) InsertAtTailOptimized(val int) {
	if ll.Head == nil {
		ll.InsertAtHead(val)
		return
	}
	node := NewDefaultNode(val)
	ll.Tail.Next = node
	ll.Tail = node
	ll.Size += 1
}

/*
Delete the node at its first occurance
Time O(n)
*/
func (ll *LinkedList) DeleteNode(val int) {
	curr := ll.Head
	// head is nil
	if curr == nil {
		return
	}

	// delete head
	if curr.Val == val {
		ll.Head = curr.Next
		ll.Size -= 1
		if ll.Head == nil {
			ll.Tail = nil
		}
		return
	}

	// delete middle or tail
	prev := curr
	curr = curr.Next
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			ll.Size -= 1
			// deleting tail
			if curr.Next == nil {
				ll.Tail = prev
			}
			return
		}
		prev = curr
		curr = curr.Next
	}
}

/*
Find length
time O(n)
*/
func (ll *LinkedList) FindLength() int {
	length := 0
	if ll.Head == nil {
		return length
	}

	curr := ll.Head
	for curr != nil {
		length += 1
		curr = curr.Next
	}
	return length
}

func (ll *LinkedList) PrintLinkedList() {
	curr := ll.Head
	if ll.Head == nil {
		fmt.Printf("{  ||  Head  =>  nil  Tail  =>  nil  size  =>  %d } \n", 0)
		return
	}
	fmt.Print("{ ")
	for curr != nil {
		fmt.Printf("%d  =>  ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("  ||  Head  =>  %d  Tail  =>  %d  size  =>  %d } \n", ll.Head.Val, ll.Tail.Val, ll.Size)
}

func NewDefaultLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Size: 0,
		Tail: nil,
	}
}
