package binarytree

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/trees"
	"github.com/OladapoAjala/datastructures/trees/node"
)

var left = false

type BinaryTree[T comparable] struct {
	Root   *node.Node[T]
	size   int32
	height int32
}

type IBinaryTree[T comparable] interface {
	trees.ITrees[T]
}

// var _ IBinaryTree[string] = new(BinaryTree[string])

func NewBinaryTree[T comparable]() *BinaryTree[T] {
	return new(BinaryTree[T])
}

func (bt *BinaryTree[T]) Insert(data T) error {
	if data == *new(T) {
		return fmt.Errorf("empty data")
	}
	if bt.Root == nil {
		bt.Root = node.NewNode[T](data)
		return nil
	}

	bt.insert(bt.Root, data)
	bt.size++
	return nil
}

func (bt *BinaryTree[T]) insert(n *node.Node[T], data T) {
	if n.Left == nil {
		n.Left = node.NewNode[T](data)
		return
	} else if n.Right == nil {
		n.Right = node.NewNode[T](data)
		return
	}

	left = !left
	if left {
		bt.insert(n.Left, data)
		return
	}
	bt.insert(n.Right, data)
}

func (bt *BinaryTree[T]) TraversalOrder(n *node.Node[T]) ([]T, error) {
	if bt.size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []T{}, nil
	}

	leftOrder, err := bt.TraversalOrder(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := bt.TraversalOrder(n.Right)
	if err != nil {
		return nil, err
	}

	output := append(leftOrder, n.Data)
	output = append(output, rightOrder...)
	return output, nil
}

func (bt *BinaryTree[T]) PreOrderTraversal(n *node.Node[T]) ([]T, error) {
	if bt.size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []T{}, nil
	}

	leftOrder, err := bt.PreOrderTraversal(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := bt.PreOrderTraversal(n.Right)
	if err != nil {
		return nil, err
	}

	output := []T{n.Data}
	output = append(output, leftOrder...)
	output = append(output, rightOrder...)
	return output, nil
}

func (bt *BinaryTree[T]) PostOrderTraversal(n *node.Node[T]) ([]T, error) {
	if bt.size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []T{}, nil
	}

	leftOrder, err := bt.PostOrderTraversal(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := bt.PostOrderTraversal(n.Right)
	if err != nil {
		return nil, err
	}

	output := append(leftOrder, rightOrder...)
	output = append(output, n.Data)
	return output, nil
}

func (bt *BinaryTree[T]) GetSize() int32 {
	return bt.size
}

func (bt *BinaryTree[T]) GetHeight() int32 {
	return bt.height
}
