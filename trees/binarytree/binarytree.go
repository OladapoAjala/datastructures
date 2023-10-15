package binarytree

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/trees"
	"github.com/OladapoAjala/datastructures/trees/node"
)

var left = false

type BinaryTree[T comparable] struct {
	Root   *node.Node[T]
	Size   int32
	Height int32
}

type IBinaryTree[T comparable] interface {
	trees.ITrees[T]
}

var _ IBinaryTree[string] = new(BinaryTree[string])

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
	bt.Size++
	return nil
}

func (bt *BinaryTree[T]) insert(n *node.Node[T], data T) {
	if n.Left == nil {
		n.Left = node.NewNode[T](data)
		n.Left.Parent = n
		return
	} else if n.Right == nil {
		n.Right = node.NewNode[T](data)
		n.Right.Parent = n
		return
	}

	left = !left
	if left {
		bt.insert(n.Left, data)
		return
	}
	bt.insert(n.Right, data)
}

func (bt *BinaryTree[T]) Delete(n *node.Node[T]) error {
	if n.IsLeaf() {
		if n.Parent.Left == n {
			n.Parent.Left = nil
		} else {
			n.Parent.Right = nil
		}
		return nil
	}

	if n.Left != nil {
		pre, err := bt.Predecessor(n)
		if err != nil {
			return err
		}
		tmp := pre.Data
		pre.Data = n.Data
		n.Data = tmp
		return bt.Delete(pre)
	}

	suc, err := bt.Successor(n)
	if err != nil {
		return err
	}
	tmp := suc.Data
	suc.Data = n.Data
	n.Data = tmp
	return bt.Delete(suc)
}

func (bt *BinaryTree[T]) InsertAfter(n *node.Node[T], data T) error {
	if n == nil {
		return fmt.Errorf("empty node")
	}

	newNode := node.NewNode[T](data)
	if bt.Root == nil {
		bt.Root = newNode
		bt.Size++
		return nil
	}

	if n.Right == nil {
		n.Right = newNode
		newNode.Parent = n
		bt.Size++
		return nil
	}

	successor, err := bt.Successor(n)
	if err != nil {
		return err
	}
	successor.Left = newNode
	newNode.Parent = successor
	bt.Size++
	return nil
}

func (bt *BinaryTree[T]) InsertBefore(n *node.Node[T], data T) error {
	newNode := node.NewNode[T](data)
	if bt.Root == nil {
		bt.Root = newNode
		bt.Size++
		return nil
	}

	if n.Left == nil {
		n.Left = newNode
		newNode.Parent = n
		bt.Size++
		return nil
	}

	predecessor, err := bt.Predecessor(n)
	if err != nil {
		return err
	}
	predecessor.Right = newNode
	newNode.Parent = predecessor
	bt.Size++
	return nil
}

func (bt *BinaryTree[T]) SubTreeFirst(n *node.Node[T]) (*node.Node[T], error) {
	if bt.Size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Left == nil {
		return n, nil
	}
	return bt.SubTreeFirst(n.Left)
}

func (bt *BinaryTree[T]) SubTreeLast(n *node.Node[T]) (*node.Node[T], error) {
	if bt.Size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Right == nil {
		return n, nil
	}
	return bt.SubTreeLast(n.Right)
}

func (bt *BinaryTree[T]) Successor(n *node.Node[T]) (*node.Node[T], error) {
	if bt.Size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Right == nil {
		return bt.climbLeft(n)
	}
	return bt.SubTreeFirst(n.Right)
}

func (bt *BinaryTree[T]) climbLeft(n *node.Node[T]) (*node.Node[T], error) {
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}
	if n.Parent == nil {
		return nil, fmt.Errorf("node %v has no parent", n)
	}

	if n == n.Parent.Left {
		return n.Parent, nil
	}
	return bt.climbLeft(n.Parent)
}

func (bt *BinaryTree[T]) Predecessor(n *node.Node[T]) (*node.Node[T], error) {
	if bt.Size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Left == nil {
		return bt.climbRight(n)
	}
	return bt.SubTreeLast(n.Left)
}

func (bt *BinaryTree[T]) climbRight(n *node.Node[T]) (*node.Node[T], error) {
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}
	if n.Parent == nil {
		return nil, fmt.Errorf("node %v has no parent", n)
	}

	if n == n.Parent.Right {
		return n.Parent, nil
	}
	return bt.climbRight(n.Parent)
}

func (bt *BinaryTree[T]) TraversalOrder(n *node.Node[T]) ([]T, error) {
	if bt.Size < 1 {
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
	if bt.Size < 1 {
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
	if bt.Size < 1 {
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
	return bt.Size
}

func (bt *BinaryTree[T]) GetHeight() int32 {
	return bt.Height
}
