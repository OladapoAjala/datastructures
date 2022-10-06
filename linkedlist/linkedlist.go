package linkedlist

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/node"
)

type LinkedList[T any] struct {
	Length int32
	Head   *node.Node[T]
	Tail   *node.Node[T]
}

type ILinkedList[T any] interface {
	Add(T) error
	AddFirst(T) error
	Clear() error
	GetNode(int32) (*node.Node[T], error)
	Insert(int32, T) error
	IsEmpty() bool
	Remove(int32) (T, error)
	// Reverse() *LinkedList[T]
	Size() int32
	// ToArray() []T
}

var _ ILinkedList[string] = new(LinkedList[string])

func NewList[T any](data ...T) *LinkedList[T] {
	list := new(LinkedList[T])
	for _, d := range data {
		list.Add(d)
	}
	return list
}

func (l *LinkedList[T]) Add(data T) error {
	if l.IsEmpty() {
		newNode := node.NewNode[T]()
		newNode.Data = data
		l.Head, l.Tail = newNode, newNode
		l.Length++
		return nil
	}

	newNode := node.NewNode[T]()
	newNode.Data = data
	newNode.Prev = l.Tail
	newNode.Next = nil

	l.Tail.Next = newNode
	l.Tail = newNode
	l.Length++

	return nil
}

func (l *LinkedList[T]) AddFirst(data T) error {
	if l.IsEmpty() {
		newNode := node.NewNode[T]()
		newNode.Data = data
		l.Head, l.Tail = newNode, newNode
		l.Length++
		return nil
	}

	newNode := node.NewNode[T]()
	newNode.Data = data
	newNode.Next = l.Head
	newNode.Prev = nil

	l.Head.Prev = newNode
	l.Head = newNode
	l.Length++

	return nil
}

func (l *LinkedList[T]) Clear() error {
	var zero T
	for it := l.Head; it != nil; it = it.Next {
		it.Data = zero
	}

	return nil
}

func (l *LinkedList[T]) GetNode(index int32) (*node.Node[T], error) {
	for it := l.Head; it != nil; it = it.Next {
		if index == 0 {
			return it, nil
		}
		index--
	}

	return nil, fmt.Errorf("node not found")
}

func (l *LinkedList[T]) Insert(index int32, data T) error {
	if l.IsEmpty() {
		newNode := node.NewNode[T]()
		newNode.Data = data
		l.Head, l.Tail = newNode, newNode
		l.Length++
		return nil
	}

	oldNode, err := l.GetNode(index)
	if err != nil {
		return fmt.Errorf("insertion failed: %v", err)
	}

	newNode := node.NewNode[T]()
	newNode.Data = data
	newNode.Prev = oldNode.Prev
	newNode.Next = oldNode

	oldNode.Prev.Next = newNode
	oldNode.Prev = newNode
	l.Length++

	return nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Length == 0
}

func (l *LinkedList[T]) Remove(index int32) (T, error) {
	var zero T
	if l.IsEmpty() {
		return zero, fmt.Errorf("cannot remove from empty list")
	}

	oldNode, err := l.GetNode(index)
	if err != nil {
		return zero, err
	}

	if l.Size() == 1 {
		l.Tail = nil
		l.Head = nil
		l.Length--
		return oldNode.Data, nil
	}

	if l.Tail == oldNode {
		l.Tail.Prev.Next = nil
		l.Tail = l.Tail.Prev
		l.Length--
		return oldNode.Data, nil
	}

	if l.Head == oldNode {
		l.Head = l.Head.Next
		l.Head.Prev = nil
		l.Length--
		return oldNode.Data, nil
	}

	oldNode.Next.Prev = oldNode.Prev
	oldNode.Prev.Next = oldNode.Next
	l.Length--
	return oldNode.Data, nil
}

func (l *LinkedList[T]) Size() int32 {
	return l.Length
}
