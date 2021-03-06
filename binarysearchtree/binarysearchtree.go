package binarysearchtree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

const GREATER int32 = 1
const LESSER int32 = 0

type node[T constraints.Ordered] struct {
	data  T
	left  *node[T]
	right *node[T]
}

func NewNode[T constraints.Ordered](data T) *node[T] {
	return &node[T]{
		data: data,
	}
}

type BinarySearchTree[T constraints.Ordered] struct {
	nodeCount int32
	Root      *node[T]
}

type IBinarySearchTree[T constraints.Ordered] interface {
	Add(T) error
	add(*node[T], T) *node[T]
	Contains(T) bool
	Remove(T) error
	remove(*node[T], T)
	Size() int32
}

// var _ IBinarySearchTree[string] = new(BinarySearchTree[string])

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return new(BinarySearchTree[T])
}

func (bst *BinarySearchTree[T]) Add(data T) error {
	if isPresent := bst.Contains(data); isPresent {
		return fmt.Errorf("%v present in tree", data)
	}

	bst.Root = add(bst.Root, data)
	bst.nodeCount++
	return nil
}

func add[T constraints.Ordered](node *node[T], data T) *node[T] {
	if node == nil {
		return NewNode(data)
	}

	if data <= node.data {
		node.left = add(node.left, data)
	} else {
		node.right = add(node.right, data)
	}

	return node
}

// func compare[T constraints.Ordered](a, b T) int32 {
// 	if a >= b {
// 		return GREATER
// 	}
// 	return LESSER
// }

func (bst *BinarySearchTree[T]) Contains(data T) bool {
	return contains(bst.Root, data)
}

func contains[T constraints.Ordered](node *node[T], data T) bool {
	if node == nil {
		return false
	}

	if data < node.data {
		return contains(node.left, data)
	} else if data > node.data {
		return contains(node.right, data)
	}

	return true
}

func (bst *BinarySearchTree[T]) Size() int32 {
	return bst.nodeCount
}

/***
	NON-RECURSIVE IMPLEMENTATION

	1. ADD
	func (bst *BinarySearchTree[T]) Add(data T) error {
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
			nextNode.left = NewNode(data)
			bst.nodeCount++
		}

		if data > currData {
			nextNode.right = NewNode(data)
			bst.nodeCount++
		}

		return nil
	}

	2. CONTAINS
	func (bst *BinarySearchTree[T]) Contains(data T) (bool, *node[T]) {
		if bst.nodeCount == 0 {
			return false, nil
		}
		nextNode := bst.Root
		currData := bst.Root.data
		if currData == data {
			return true, nextNode
		}

		for nextNode.left != nil && nextNode.right != nil {
			if data <= currData {
				nextNode = nextNode.left
				currData = nextNode.data
			} else {
				nextNode = nextNode.right
				currData = nextNode.data
			}

			if currData == data {
				return true, nextNode
			}
		}

		return false, nil
	}
***/
