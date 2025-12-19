package linkedlists

import "fmt"

func LinkedListProblems() {
	testingNewLinkedList()
}

func testingNewLinkedList() {
	ll := NewDeailtLinkedList()
	ll.InsertAtHead(5)
	ll.InsertAtHead(7)
	ll.InsertAtHead(22)
	ll.InsertAtTail(24)
	ll.PrintLinkedList()

	ll2 := NewDeailtLinkedList()
	ll2.InsertAtTail(55)
	ll2.PrintLinkedList()

}

// func testingNewNode() {
// 	head := NewDefaultNode(2)
// 	head.Next = NewDefaultNode(5)
// 	head.Next.Next = NewDefaultNodeWithNext(7, NewDefaultNode(22))
// 	fmt.Println(head)
// 	PrintNodes(head)
// }

func PrintNodes(head *Node) {
	curr := head
	fmt.Println()
	fmt.Print("{ ")
	for curr != nil {
		fmt.Printf("%d  =>  ", curr.Val)
		curr = curr.Next
	}
	fmt.Print("}")
	fmt.Println()
}
