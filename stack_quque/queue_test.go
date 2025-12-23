package stackquque

import "testing"

func TestQueueBasicOperations(t *testing.T) {
	myQueue := NewDefaultQueue(3)
	t.Log("Queue initialized")

	// enqueue
	myQueue.enqueue(11)
	t.Log("Enqueued 11")

	myQueue.enqueue(22)
	t.Log("Enqueued 22")

	myQueue.enqueue(33)
	t.Log("Enqueued 33")

	if myQueue.getSize() != 3 {
		t.Fatalf("expected size 3, got %d", myQueue.getSize())
	}
	t.Log("Size check passed")

	// check front
	if myQueue.front() != 11 {
		t.Fatalf("expected front 11, got %d", myQueue.front())
	}
	t.Log("Front check passed")

	// dequeue
	if myQueue.dequeue() != 11 {
		t.Fatalf("expected dequeue 11")
	}
	t.Log("Dequeued 11")

	if myQueue.dequeue() != 22 {
		t.Fatalf("expected dequeue 22")
	}
	t.Log("Dequeued 22")

	if myQueue.getSize() != 1 {
		t.Fatalf("expected size 1, got %d", myQueue.getSize())
	}
	t.Log("Size after dequeue passed")

	if myQueue.front() != 33 {
		t.Fatalf("expected front 33, got %d", myQueue.front())
	}
	t.Log("Final front check passed")
}

func TestQueueCircularBehavior(t *testing.T) {
	myQueue := NewDefaultQueue(3)
	myQueue.enqueue(1)
	myQueue.enqueue(2)
	myQueue.enqueue(3)
	t.Log("Enqueue check passed")
	if myQueue.dequeue() != 1 {
		t.Fatal("expected 1")
	}

	// should add in circular way
	myQueue.enqueue(4)

	if myQueue.curr_size != 3 {
		t.Fatalf("expected size 3, got %d", myQueue.curr_size)
	}
	t.Log("Circular add check passed")

	if myQueue.dequeue() != 2 {
		t.Fatal("expected 2")
	}

	if myQueue.dequeue() != 3 {
		t.Fatal("expected 3")
	}

	if myQueue.dequeue() != 4 {
		t.Fatal("expected 4")
	}

	if myQueue.curr_size != 0 {
		t.Fatalf("expected size 0, got %d", myQueue.curr_size)
	}
	t.Log("Circular removal check passed")

}

func TestQueueUnderflow(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic on underflow")
		}
	}()

	q := NewDefaultQueue(2)
	q.dequeue() // should panic
	t.Log("Queue underflow passed")
}

func TestQueueOverflow(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic on overflow")
		}
	}()

	q := NewDefaultQueue(2)
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3) // should panic
	t.Log("Overflow panic passed")
}
