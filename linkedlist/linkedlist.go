package linkedlist

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/node"
)

type LinkedList struct {
	Length int32
	Head   *node.Node
	Tail   *node.Node
}

type ILinkedList interface {
	Add(any) error
	AddFirst(any) error
	Clear() error
	GetNode(int32) (*node.Node, error)
	Insert(int32, any) error
	IsEmpty() bool
	Remove(int32) (any, error)
	// Reverse() *LinkedList
	Size() int32
	// ToArray() []any
}

var _ ILinkedList = new(LinkedList)

func NewList(data ...any) *LinkedList {
	list := new(LinkedList)
	for _, d := range data {
		list.Add(d)
	}
	return list
}

func (l *LinkedList) Add(data any) error {
	if l.IsEmpty() {
		newNode := node.NewNode()
		newNode.Data = data
		l.Head, l.Tail = newNode, newNode
		l.Length++
		return nil
	}

	newNode := node.NewNode()
	newNode.Data = data
	newNode.Prev = l.Tail
	newNode.Next = nil

	l.Tail.Next = newNode
	l.Tail = newNode
	l.Length++

	return nil
}

func (l *LinkedList) AddFirst(data any) error {
	if l.IsEmpty() {
		newNode := node.NewNode()
		newNode.Data = data
		l.Head, l.Tail = newNode, newNode
		l.Length++
		return nil
	}

	newNode := node.NewNode()
	newNode.Data = data
	newNode.Next = l.Head
	newNode.Prev = nil

	l.Head.Prev = newNode
	l.Head = newNode
	l.Length++

	return nil
}

func (l *LinkedList) Clear() error {
	for it := l.Head; it != nil; it = it.Next {
		it.Data = nil
	}

	return nil
}

func (l *LinkedList) GetNode(index int32) (*node.Node, error) {
	for it := l.Head; it != nil; it = it.Next {
		if index == 0 {
			return it, nil
		}
		index--
	}

	return nil, fmt.Errorf("node not found")
}

func (l *LinkedList) Insert(index int32, data any) error {
	if l.IsEmpty() {
		newNode := node.NewNode()
		newNode.Data = data
		l.Head, l.Tail = newNode, newNode
		l.Length++
		return nil
	}

	oldNode, err := l.GetNode(index)
	if err != nil {
		return fmt.Errorf("insertion failed: %v", err)
	}

	newNode := node.NewNode()
	newNode.Data = data
	newNode.Prev = oldNode.Prev
	newNode.Next = oldNode

	oldNode.Prev.Next = newNode
	oldNode.Prev = newNode
	l.Length++

	return nil
}

func (l *LinkedList) IsEmpty() bool {
	return l.Length == 0
}

func (l *LinkedList) Remove(index int32) (any, error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("cannot remove from empty list")
	}

	oldNode, err := l.GetNode(index)
	if err != nil {
		return nil, err
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

func (l *LinkedList) Size() int32 {
	return l.Length
}
