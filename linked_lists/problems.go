package linkedlists

import "fmt"

func LinkedListProblems() {
	testingNewLinkedList()
}

func testingNewLinkedList() {
	ll := NewDefaultLinkedList()
	ll.PrintLinkedList()
	ll.InsertAtHead(5)
	ll.InsertAtHead(7)
	ll.InsertAtHead(22)
	ll.InsertAtTailOptimized(24)
	ll.InsertAtTailOptimized(25)
	ll.InsertAtTailOptimized(26)
	ll.InsertAtTailOptimized(27)
	ll.InsertAtHead(11)
	fmt.Println("before deleting in middle")
	ll.PrintLinkedList()
	fmt.Println("after deleting first")
	ll.DeleteNode(11)
	ll.PrintLinkedList()
	ll.DeleteNode(25)
	fmt.Println("after deleting 25")
	ll.PrintLinkedList()
	fmt.Println("deleting tail 27")
	ll.DeleteNode(27)
	ll.PrintLinkedList()

	fmt.Println()
	ll.InsertAt(55, 1)
	fmt.Println("after adding at 1st")
	ll.PrintLinkedList()
	fmt.Println("inserting 44 at 4th")
	ll.InsertAt(44, 4)
	ll.PrintLinkedList()
	fmt.Println("adding at second last")
	ll.InsertAt(66, 6)
	ll.PrintLinkedList()
	fmt.Println("adding at last")
	ll.InsertAt(88, 8)
	ll.PrintLinkedList()
	fmt.Printf("Length of ll %d \n", ll.FindLength())
	ll3 := NewDefaultLinkedList()
	fmt.Printf("Length of ll %d \n", ll3.FindLength())
	ll3.InsertAtTail(5)
	fmt.Printf("Length of ll %d \n", ll3.FindLength())

	ll2 := NewDefaultLinkedList()
	ll2.InsertAtTailOptimized(55)
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
