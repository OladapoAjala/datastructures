package linkedlist

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/node"
	"github.com/OladapoAjala/datastructures/sequences"
)

type LinkedList[T comparable] struct {
	length int32
	Head   *node.Node[T]
	Tail   *node.Node[T]
}

type ILinkedList[T comparable] interface {
	sequences.Sequencer[T]
	GetNode(int32) (*node.Node[T], error)
	ToArray() ([]T, error)
	Reverse() error
	Clear() error
}

var _ ILinkedList[string] = new(LinkedList[string])

func NewList[T comparable](data ...T) *LinkedList[T] {
	list := new(LinkedList[T])
	for _, d := range data {
		list.InsertLast(d)
	}
	return list
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

func (l *LinkedList[T]) GetData(index int32) (T, error) {
	node, err := l.GetNode(index)
	if err != nil {
		return *new(T), fmt.Errorf("data not found")

	}
	return node.Data, nil
}

func (l *LinkedList[T]) Contains(data T) bool {
	for it := l.Head; it != nil; it = it.Next {
		if it.Data == data {
			return true
		}
	}

	return false
}

func (l *LinkedList[T]) InsertFirst(data T) error {
	if l.IsEmpty() {
		newNode := node.NewNode[T]()
		newNode.Data = data
		l.Head, l.Tail = newNode, newNode
		l.length++
		return nil
	}

	newNode := node.NewNode[T]()
	newNode.Data = data
	newNode.Next = l.Head
	newNode.Prev = nil

	l.Head.Prev = newNode
	l.Head = newNode
	l.length++

	return nil
}

func (l *LinkedList[T]) InsertLast(data T) error {
	if l.IsEmpty() {
		return l.InsertFirst(data)
	}

	newNode := node.NewNode[T]()
	newNode.Data = data
	newNode.Prev = l.Tail
	newNode.Next = nil

	l.Tail.Next = newNode
	l.Tail = newNode
	l.length++

	return nil
}

func (l *LinkedList[T]) Insert(index int32, data T) error {
	if l.IsEmpty() {
		return l.InsertFirst(data)
	}

	if index >= l.Size() {
		return l.InsertLast(data)
	}

	oldNode, err := l.GetNode(index)
	if err != nil {
		return fmt.Errorf("insertion failed: %v", err)
	}

	oldNode.Data = data
	return nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.length == 0
}

func (l *LinkedList[T]) Delete(index int32) error {
	if l.IsEmpty() {
		return fmt.Errorf("cannot remove from empty list")
	}

	oldNode, err := l.GetNode(index)
	if err != nil {
		return err
	}

	if l.Size() == 1 {
		l.Tail = nil
		l.Head = nil
		l.length--
		return nil
	}

	if l.Tail == oldNode {
		l.Tail.Prev.Next = nil
		l.Tail = l.Tail.Prev
		l.length--
		return nil
	}

	if l.Head == oldNode {
		l.Head = l.Head.Next
		l.Head.Prev = nil
		l.length--
		return nil
	}

	oldNode.Next.Prev = oldNode.Prev
	oldNode.Prev.Next = oldNode.Next
	l.length--
	return nil
}

func (l *LinkedList[T]) DeleteFirst() error {
	if l.IsEmpty() {
		return fmt.Errorf("cannot remove from empty list")
	}

	if l.Size() == 1 {
		l.Tail = nil
		l.Head = nil
		l.length--
		return nil
	}

	l.Head = l.Head.Next
	l.Head.Prev = nil
	l.length--
	return nil
}

func (l *LinkedList[T]) DeleteLast() error {
	if l.IsEmpty() {
		return fmt.Errorf("cannot remove from empty list")
	}

	if l.Size() == 1 {
		l.Tail = nil
		l.Head = nil
		l.length--
		return nil
	}

	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
	l.length--
	return nil
}

func (l *LinkedList[T]) Clear() error {
	var zero T
	for it := l.Head; it != nil; it = it.Next {
		it.Data = zero
	}

	return nil
}

func (l *LinkedList[T]) Reverse() error {
	if l.IsEmpty() {
		return fmt.Errorf("cannot reverse empty list")
	}

	for it := l.Head; it != nil; it = it.Prev {
		next := it.Next
		it.Next = it.Prev
		it.Prev = next
	}

	head := l.Head
	l.Head = l.Tail
	l.Tail = head
	return nil
}

func (l *LinkedList[T]) Size() int32 {
	return l.length
}

func (l *LinkedList[T]) ToArray() ([]T, error) {
	array := make([]T, l.Size())
	i := 0
	for it := l.Head; it != nil; it = it.Next {
		array[i] = it.Data
		i++
	}
	return array, nil
}
