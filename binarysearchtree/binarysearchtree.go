package binarysearchtree

import (
	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	data  T
	left  *node[T]
	right *node[T]
}

type BinarySearchTree[T constraints.Ordered] struct {
	nodeCount int32
	Root      *node[T]
}

type IBinarySearchTree[T constraints.Ordered] interface {
	Insert(T) error
}

var _ IBinarySearchTree[string] = new(BinarySearchTree[string])

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		Root: new(node[T]),
	}
}

func (bst *BinarySearchTree[T]) Insert(data T) error {
	if bst.nodeCount == 0 {
		bst.Root.data = data
		bst.nodeCount++
		return nil
	}

	currData := bst.Root.data
	nextNode := bst.Root

	for nextNode.left != nil && nextNode.right != nil {
		if data <= currData {
			nextNode = nextNode.left
			currData = nextNode.data
			continue
		}

		nextNode = nextNode.right
		currData = nextNode.data
	}

	if data <= currData {
		nextNode.left = new(node[T])
		nextNode.left.data = data
		bst.nodeCount++
	}

	if data > currData {
		nextNode.right = new(node[T])
		nextNode.right.data = data
		bst.nodeCount++
	}

	return nil
}
