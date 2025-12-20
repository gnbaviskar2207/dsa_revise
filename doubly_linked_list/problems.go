package doublylinkedlist

import "fmt"

func DoublyLinkedProblems() {
	dll := NewDefaultDoublyLinkedListEmpty()
	dll.insertAtTail(55)
	dll.insertHead(5)
	dll.insertAtTail(555)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()
	fmt.Print("after adding 22 at 2\n")
	dll.insertAt(2, 22)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()

	fmt.Print("after adding 44 at 4\n")
	dll.insertAt(4, 44)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()

	fmt.Print("after adding 66 at 6 which is tail\n")
	dll.insertAt(6, 66)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()

	fmt.Print("deleting at 1st\n")
	dll.deleteAt(1)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()

	fmt.Print("deleting at 0th\n")
	dll.deleteAt(0)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()

	fmt.Print("deleting at 6th more than len\n")
	dll.deleteAt(6)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()

	fmt.Print("deleting at tail which is 5th\n")
	dll.deleteAt(5)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()

	fmt.Print("deleting 3rd which is 44\n")
	dll.deleteAt(3)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()

	fmt.Print("adding 5 elements\n")
	dll.insertAtTail(5)
	dll.insertAtTail(6)
	dll.insertAtTail(7)
	dll.insertAtTail(8)
	dll.insertAtTail(9)
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()

	fmt.Print("reverse the dll\n")
	dll.reverse()
	dll.printDoublyLinkedList()
	dll.printDoublyLinkedListReverse()
}
