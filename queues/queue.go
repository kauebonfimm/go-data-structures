package queue

import "fmt"

type IQueue interface {
	Enqueue(string) error
	Dequeue() (string, error)
	Peek() (string, error)
	Lenght() int
	Clear() error
	IsEmpty() bool
	IsFull() bool
}

type Queue struct {
	queue  []string
	lenght int
}

func NewQueue(lenght int) IQueue {
	return &Queue{
		queue:  make([]string, 0, lenght),
		lenght: lenght,
	}
}

func (q *Queue) Enqueue(data string) error {
	if q.IsFull() {
		return fmt.Errorf("error to enqueue: the queue is full")
	}

	q.queue = append(q.queue, data)

	return nil
}

func (q *Queue) Dequeue() (string, error) {
	if q.IsEmpty() {
		return "", fmt.Errorf("error to dequeue: the queue is empty")
	}

	data := q.queue[0]

	q.queue = q.queue[1:]

	return data, nil
}

func (q *Queue) Peek() (string, error) {
	if q.IsEmpty() {
		return "", fmt.Errorf("error to dequeue: the queue is empty")
	}

	return q.queue[0], nil
}

func (q *Queue) Lenght() int {
	return len(q.queue)
}

func (q *Queue) Clear() error {
	q.queue = make([]string, 0, q.lenght)
	return nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue) IsFull() bool {
	return len(q.queue) == q.lenght
}
