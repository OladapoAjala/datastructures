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

func NewList(data ...interface{}) *LinkedList {
	list := new(LinkedList)
	for _, d := range data {
		list.Add(d)
	}
	return list
}

func (l *LinkedList) Add(data interface{}) (bool, error) {
	if l.Head == nil {
		newNode := node.NewNode()
		newNode.Data = data

		l.Head, l.Tail = newNode, newNode
		l.Length++

		return true, nil
	}

	if l.Contains(data) {
		return false, fmt.Errorf("List already contains node %v", data)
	}

	newNode := node.NewNode()
	newNode.Data = data
	newNode.Index = l.Length
	newNode.Prev = l.Tail
	newNode.Next = nil

	l.Tail.Next = newNode
	l.Tail = newNode
	l.Length++

	return true, nil
}

func (l *LinkedList) Prepend(data interface{}) (bool, error) {
	if l.Head == nil {
		newNode := node.NewNode()
		newNode.Data = data

		l.Head, l.Tail = newNode, newNode
		l.Length++

		return true, nil
	}

	if l.Contains(data) {
		return false, fmt.Errorf("List already contains node %v", data)
	}

	newNode := node.NewNode()
	newNode.Data = data
	newNode.Index = l.Length
	newNode.Next = l.Head
	newNode.Prev = nil

	l.ShiftRight(0)

	l.Head.Prev = newNode
	l.Head = newNode
	l.Length++
}

func (L *LinkedList) Insert(index int32, value string) {
	if L.head == nil {
		newTail := &Node{
			value: value,
			index: 0,
			prev:  nil,
			next:  nil,
		}

		L.head, L.tail = newTail, newTail
		L.length++

		return
	}

	// Get the current Node at the desired index.
	var oldNode *Node
	for it := L.head; it != nil; it = it.next {
		if it.index == index {
			oldNode = it
			break
		}
	}

	// Shift the indices of all the elements from the desired index.
	for it := oldNode; it != nil; it = it.next {
		it.index = it.index + 1
	}

	newNode := &Node{
		value: value,
		index: index,
		prev:  oldNode.prev,
		next:  oldNode,
	}

	oldNode.prev.next = newNode
	oldNode.prev = newNode
	L.length++
}

func (L *LinkedList) Remove(index int32) {
	var oldNode *Node
	for it := L.head; it != nil; it = it.next {
		if it.index == index {
			oldNode = it
			break
		}
	}

	// Re-arrange the indices
	for it := oldNode.next; it != nil; it = it.next {
		it.index--
	}

	if oldNode.next != nil {
		oldNode.next.prev = oldNode.prev
	}
	if oldNode.prev != nil {
		oldNode.prev.next = oldNode.next
	}
	L.length--
}

func (l *LinkedList) Contains(data interface{}) bool {

}

func (l *LinkedList) GetNode(index int32) (*node.Node, error) {
	for it := l.Head; it != nil; it = it.Next {
		if it.Index == index {
			return it, nil
		}
	}

	return nil, fmt.Errorf("Node not found")
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
