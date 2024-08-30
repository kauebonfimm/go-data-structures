package stack

import "fmt"

type IStack interface {
	Push(string) error
	Pop() (string, error)
	Peek() (string, error)
	Lenght() int
	Clear() error
	IsEmpty() bool
	IsFull() bool
}

func NewStack(size int) IStack {
	return &Stack{
		stack: make([]string, 0, size),
		len:   size,
	}
}

type Stack struct {
	stack []string
	len   int
}

func (s *Stack) Push(text string) error {
	if s.IsFull() {
		return fmt.Errorf("Stack is full")
	}

	s.stack = append(s.stack, text)
	return nil
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("Stack is Empty")
	}
	lastIndice := len(s.stack) - 1
	value := s.stack[lastIndice]
	s.stack = s.stack[:lastIndice]
	return value, nil
}

func (s *Stack) Peek() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("Stack is Empty")
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *Stack) Lenght() int {
	return len(s.stack)
}

func (s *Stack) Clear() error {
	s.stack = make([]string, 0, s.len)
	return nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) IsFull() bool {
	return !(len(s.stack) < s.len)
}
