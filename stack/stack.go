package stack

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/linkedlist"
)

type Stack[T any] struct {
	*linkedlist.LinkedList[T]
}

type IStack[T any] interface {
	Peek() (T, error)
	Pop() (T, error)
	Push(T) error
}

var _ IStack[int] = new(Stack[int])

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		new(linkedlist.LinkedList[T]),
	}
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, fmt.Errorf("cannot peek empty stack")
	}
	return s.Tail.Data, nil
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, fmt.Errorf("cannot pop from empty stack")
	}

	node, err := s.GetNode(s.Length - 1)
	if err != nil {
		return zero, fmt.Errorf("error getting head of stack")
	}

	_, err = s.Remove(s.Length - 1)
	if err != nil {
		return zero, fmt.Errorf("unable to delete last element in stack")
	}

	return node.Data, nil
}

func (s *Stack[T]) Push(data T) error {
	err := s.Add(data)
	if err != nil {
		return err
	}

	return nil
}
