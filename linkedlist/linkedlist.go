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
	Add(interface{}) error
	AddFirst(interface{}) error
	Clear() error
	Contains(interface{}) (bool, int32)
	GetNode(int32) (*node.Node, error)
	Insert(int32, interface{}) error
	Remove(int32) error
	ShiftLeft(int32) error
	ShiftRight(int32) error
	Size() int32
	// ToArray()
}

var _ ILinkedList = new(LinkedList)

func NewList(data ...interface{}) *LinkedList {
	list := new(LinkedList)
	for _, d := range data {
		list.Add(d)
	}
	return list
}

func (l *LinkedList) Add(data interface{}) error {
	if l.Head == nil {
		newNode := node.NewNode()
		newNode.Data = data

		l.Head, l.Tail = newNode, newNode
		l.Length++

		return nil
	}

	// if l.Contains(data) {
	// 	return false, fmt.Errorf("List already contains node %v", data)
	// }

	newNode := node.NewNode()
	newNode.Data = data
	newNode.Index = l.Length
	newNode.Prev = l.Tail
	newNode.Next = nil

	l.Tail.Next = newNode
	l.Tail = newNode
	l.Length++

	return nil
}

func (l *LinkedList) AddFirst(data interface{}) error {
	if l.Head == nil {
		newNode := node.NewNode()
		newNode.Data = data

		l.Head, l.Tail = newNode, newNode
		l.Length++

		return nil
	}

	// if l.Contains(data) {
	// 	return false, fmt.Errorf("List already contains node %v", data)
	// }

	newNode := node.NewNode()
	newNode.Data = data
	newNode.Index = l.Length
	newNode.Next = l.Head
	newNode.Prev = nil

	l.ShiftRight(0)

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

func (l *LinkedList) Contains(data interface{}) (bool, int32) {
	for it := l.Head; it != nil; it = it.Next {
		if it.Data == data {
			return true, it.Index
		}
	}

	return false, -1
}

func (l *LinkedList) GetNode(index int32) (*node.Node, error) {
	for it := l.Head; it != nil; it = it.Next {
		if it.Index == index {
			return it, nil
		}
	}

	return nil, fmt.Errorf("Node not found")
}

func (l *LinkedList) Insert(index int32, data interface{}) error {
	if l.Head == nil {
		newNode := node.NewNode()
		newNode.Data = data

		l.Head, l.Tail = newNode, newNode
		l.Length++

		return nil
	}

	err := l.ShiftRight(index)
	if err != nil {
		return fmt.Errorf("Insertion failed: %v", err)
	}

	oldNode, err := l.GetNode(index)
	if err != nil {
		return fmt.Errorf("Insertion failed: %v", err)
	}

	newNode := node.NewNode()
	newNode.Data = data
	newNode.Index = index
	newNode.Prev = oldNode.Prev
	newNode.Next = oldNode

	oldNode.Prev.Next = newNode
	oldNode.Prev = newNode
	l.Length++

	return nil
}

// Update this!
func (l *LinkedList) Remove(index int32) error {
	oldNode, err := l.GetNode(index)
	if err != nil {
		return fmt.Errorf("error")
	}

	err = l.ShiftLeft(index)
	if err != nil {
		return fmt.Errorf("error")
	}

	if oldNode.Next != nil {
		oldNode.Next.Prev = oldNode.Prev
	}

	if oldNode.Prev != nil {
		oldNode.Prev.Next = oldNode.Next
	}

	l.Length--
	return nil
}

func (l *LinkedList) ShiftLeft(index int32) error {
	n, err := l.GetNode(index)
	if err != nil {
		return err
	}

	for it := n; it != nil; it = it.Next {
		it.Index--
	}

	return nil
}

func (l *LinkedList) ShiftRight(index int32) error {
	n, err := l.GetNode(index)
	if err != nil {
		return err
	}

	for it := n; it != nil; it = it.Next {
		n.Index++
	}

	return nil
}

func (l *LinkedList) Size() int32 {
	return l.Length
}
