package stack

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/linkedlist"
)

type Stack struct {
	*linkedlist.LinkedList
}

type IStack interface {
	Peek() (any, error)
	Pop() (any, error)
	Push(any) error
}

var _ IStack = new(Stack)

func NewStack() *Stack {
	return &Stack{
		new(linkedlist.LinkedList),
	}
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("cannot peek empty stack")
	}
	return s.Tail.Data, nil
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("cannot pop from empty stack")
	}

	node, err := s.GetNode(s.Length - 1)
	if err != nil {
		return nil, fmt.Errorf("error getting head of stack")
	}

	_, err = s.Remove(s.Length - 1)
	if err != nil {
		return nil, fmt.Errorf("unable to delete last element in stack")
	}

	return node.Data, nil
}

func (s *Stack) Push(data any) error {
	err := s.Add(data)
	if err != nil {
		return err
	}

	return nil
}
