package stackquque

import "fmt"

func stackProblems() {
	myStack := NewDefaultStack(10)
	myStack.push(1)
	myStack.push(12)
	myStack.push(13)
	myStack.push(14)
	myStack.push(15)
	myStack.push(16)
	myStack.push(17)
	myStack.push(18)
	myStack.push(19)
	myStack.push(11)
	fmt.Println(myStack.top())
	myStack.pop()
	fmt.Println(myStack.pop())
	fmt.Println(myStack.top())
}

func queueProblems() {
	myQueue := NewDefaultQueue(4)
	myQueue.enqueue(1)
	myQueue.enqueue(2)
	myQueue.enqueue(3)
	myQueue.enqueue(4)
	fmt.Println(myQueue.getSize())
	// checking whether it shoots error
	myQueue.enqueue(66)
}

func StackQueueProblems() {
	// stackProblems()
	queueProblems()
}
