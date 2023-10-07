package node

import (
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Data  T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T constraints.Ordered](data T) *Node[T] {
	return &Node[T]{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (n *Node[T]) GetData() T {
	return n.Data
}

func (n *Node[T]) GetLeft() *Node[T] {
	return n.Left
}

func (n *Node[T]) GetRight() *Node[T] {
	return n.Right
}
