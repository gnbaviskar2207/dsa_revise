package doublylinkedlist

import "fmt"

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (dll *DoublyLinkedList) insertHead(val int) {
	node := NewDefaultNode(val)
	if dll.Head == nil {

		dll.Head = node
		dll.Tail = node
		dll.Size = 1
		return
	}
	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node
	dll.Size += 1
}

func (dll *DoublyLinkedList) insertAtTail(val int) {
	if dll.Tail == nil {
		dll.insertHead(val)
		return
	}
	node := NewDefaultNode(val)
	node.Prev = dll.Tail
	dll.Tail.Next = node
	dll.Tail = node
	dll.Size += 1
}

func (dll *DoublyLinkedList) printDoublyLinkedList() {
	fmt.Printf("print ==> { ")
	curr := dll.Head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("} head %d tail %d size %d \n", dll.Head.Val, dll.Tail.Val, dll.Size)
}

func (dll *DoublyLinkedList) printDoublyLinkedListReverse() {
	fmt.Printf("print <== { ")
	curr := dll.Tail
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Prev
	}
	fmt.Printf("} head %d tail %d size %d \n", dll.Head.Val, dll.Tail.Val, dll.Size)
}

func (dll *DoublyLinkedList) insertAt(index int, val int) {
	if index == 1 {
		dll.insertHead(val)
		return
	}
	if index == dll.Size+1 {
		dll.insertAtTail(val)
		return
	}
	counter := 0
	curr := dll.Head
	for curr != nil {
		counter += 1
		if counter == index {
			node := NewDefaultNode(val)
			prev := curr.Prev
			prev.Next = node
			node.Prev = prev
			node.Next = curr
			curr.Prev = node
			break
		}
		curr = curr.Next
	}
	dll.Size += 1
}

func (dll *DoublyLinkedList) deleteAt(index int) {
	if index <= 0 || index > dll.Size {
		return
	}
	if index == 1 {
		if dll.Head == nil {
			return
		}
		if dll.Head.Next == nil {
			dll.Head = nil
			dll.Tail = nil
			dll.Size -= 1
			return
		}
		nextNode := dll.Head.Next
		dll.Head.Next = nil
		nextNode.Prev = nil
		dll.Head = nextNode
		dll.Size -= 1
		return
	}
	if index == dll.Size {
		prev := dll.Tail.Prev
		dll.Tail.Prev = nil
		prev.Next = nil
		dll.Tail = prev
		dll.Size -= 1
		return
	}

	counter := 0
	curr := dll.Head
	for curr != nil {
		counter += 1
		if counter == index {
			prev := curr.Prev
			next := curr.Next
			prev.Next = next
			next.Prev = prev
			curr.Prev = nil
			curr.Next = nil
			dll.Size -= 1
			break
		}
		curr = curr.Next
	}
}

func (dll *DoublyLinkedList) reverse() {
	// Case 1: empty list or single node → nothing to reverse
	if dll.Head == nil || dll.Head.Next == nil {
		return
	}

	curr := dll.Head
	var temp *Node = nil

	// Traverse through the list
	for curr != nil {
		// Store previous pointer temporarily
		temp = curr.Prev

		// Swap Prev and Next
		curr.Prev = curr.Next
		curr.Next = temp

		// Move to "next" node
		// (which is curr.Prev after swapping)
		curr = curr.Prev
	}

	// After loop, temp will be at the old head's previous
	// Swap head and tail
	dll.Tail = dll.Head
	dll.Head = temp.Prev
}

func NewDefaultDoublyLinkedList(val int) *DoublyLinkedList {
	node := NewDefaultNode(val)
	return &DoublyLinkedList{
		Head: node,
		Tail: node,
		Size: 1,
	}
}

func NewDefaultDoublyLinkedListEmpty() *DoublyLinkedList {
	return &DoublyLinkedList{}
}
