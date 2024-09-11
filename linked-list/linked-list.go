package linkedlist

type LinkedList[T any] struct {
	Head *Node[T]
}

type Node[T any] struct {
	Data T
	Next *Node[T]
}

func (l *LinkedList[T]) Size() (counter int) {
	if l.Head == nil {
		return
	}

	current := l.Head
	for current != nil {
		counter += 1
		current = current.Next
	}
	return
}

func (l *LinkedList[T]) InsertAtBeginning(data T) {
	node := &Node[T]{
		Data: data,
		Next: l.Head,
	}
	l.Head = node
}

func (l *LinkedList[T]) RemoveAtBeginning() {
	if l.Head == nil {
		return
	}

	l.Head = l.Head.Next
}

func (l *LinkedList[T]) InsertAtLast(data T) {
	if l.Head == nil {
		l.Head = &Node[T]{
			Data: data,
			Next: nil,
		}
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node[T]{
		Data: data,
		Next: nil,
	}
}
