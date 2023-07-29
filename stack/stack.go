package stack

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
)

type Stack[T comparable] struct {
	*linkedlist.LinkedList[T]
}

type IStack[T comparable] interface {
	Peek() (T, error)
	Pop() (T, error)
	Push(T) error
}

var _ IStack[int] = new(Stack[int])

func NewStack[T comparable]() *Stack[T] {
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

	node, err := s.GetNode(s.Size() - 1)
	if err != nil {
		return zero, fmt.Errorf("error getting head of stack")
	}

	err = s.Delete(s.Size() - 1)
	if err != nil {
		return zero, fmt.Errorf("unable to delete last element in stack")
	}

	return node.Data, nil
}

func (s *Stack[T]) Push(data T) error {
	err := s.InsertLast(data)
	if err != nil {
		return err
	}

	return nil
}
