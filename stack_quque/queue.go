package stackquque

type Queue struct {
	size      int
	curr_size int
	start     int
	end       int
	arr       []int
}

func (q *Queue) enqueue(ele int) {
	if q.curr_size >= q.size {
		panic("queue overflow")
	}
	if q.curr_size == 0 {
		q.start = 0
		q.end = 0
	} else {
		q.end = (q.end + 1) % q.size
	}
	q.arr[q.end] = ele
	q.curr_size += 1
}

func (q *Queue) dequeue() int {
	if q.curr_size == 0 {
		panic("queue underflow")
	}
	ele := q.arr[q.start]
	if q.curr_size == 1 {
		q.start = -1
		q.end = -1
	} else {
		q.start = (q.start + 1) % q.size
	}
	q.curr_size -= 1
	return ele
}

func (q *Queue) front() int {
	if q.curr_size == 0 {
		panic("queue underflow")
	}
	return q.arr[q.start]
}

func (q *Queue) getSize() int {
	return q.curr_size
}

func NewDefaultQueue(size int) *Queue {
	return &Queue{
		size:      size,
		curr_size: 0,
		start:     -1,
		end:       -1,
		arr:       make([]int, size),
	}
}
