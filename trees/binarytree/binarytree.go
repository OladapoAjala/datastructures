package binarytree

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/sequences"
	"github.com/OladapoAjala/datastructures/trees"
	"github.com/OladapoAjala/datastructures/trees/node"
)

type BinaryTree[T comparable] struct {
	Root   *node.Node[T]
	Height int32
}

type IBinaryTree[T comparable] interface {
	trees.ITrees[T]
	sequences.Sequencer[T]
	InsertAfter(*node.Node[T], *node.Node[T]) error
	InsertBefore(*node.Node[T], *node.Node[T]) error
	SubTree(*node.Node[T], int32) (*node.Node[T], error)
}

// var _ IBinaryTree[string] = new(BinaryTree[string])

// SEQUENCE METHODS
func (bt *BinaryTree[T]) GetData(index int32) (T, error) {
	n, err := bt.getNode(index)
	if err != nil {
		return *new(T), nil
	}
	return n.GetData(), nil
}

func (bt *BinaryTree[T]) Contains(data T) bool {
	return bt.contains(bt.Root, data)
}

func (bt *BinaryTree[T]) contains(n *node.Node[T], data T) bool {
	if n == nil {
		return false
	}
	if n.Data == data {
		return true
	}

	left := bt.contains(n.Left, data)
	if left {
		return left
	}
	return bt.contains(n.Right, data)
}

func (bt *BinaryTree[T]) Insert(index int32, data T) error {
	if data == *new(T) {
		return fmt.Errorf("empty data")
	}
	if index > bt.GetSize() {
		return fmt.Errorf("index %d is larger than size %d", index, bt.GetSize())
	}

	newNode := node.NewNode[T](data)
	if bt.Root == nil {
		bt.Root = newNode
		return nil
	}
	if index == bt.GetSize() {
		last := bt.SubTree(bt.Root, bt.GetSize()-1)
		last.Right = newNode
		newNode.Parent = last
		bt.updateSize(last)
		return nil
	}

	sub := bt.SubTree(bt.Root, index)
	err := bt.InsertBefore(sub, newNode)
	if err != nil {
		return err
	}
	return nil
}

func (bt *BinaryTree[T]) Delete(index int32) error {
	n, err := bt.getNode(index)
	if err != nil {
		return err
	}
	return bt.delete(n)
}

func (bt *BinaryTree[T]) delete(n *node.Node[T]) error {
	if n.IsLeaf() {
		if n.Parent.Left == n {
			n.Parent.Left = nil
		} else {
			n.Parent.Right = nil
		}
		bt.updateSize(n.Parent)
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
		return bt.delete(pre)
	}

	suc, err := bt.Successor(n)
	if err != nil {
		return err
	}
	tmp := suc.Data
	suc.Data = n.Data
	n.Data = tmp
	return bt.delete(suc)
}

// TREE METHODS
func NewBinaryTree[T comparable](data ...T) *BinaryTree[T] {
	bt := new(BinaryTree[T])
	for i, item := range data {
		err := bt.Insert(int32(i), item)
		if err != nil {
			return nil
		}
	}
	return bt
}

func (bt *BinaryTree[T]) updateSize(n *node.Node[T]) {
	if n == nil {
		return
	}
	var nl, nr int32
	if n.Left != nil {
		nl = n.Left.Size
	}
	if n.Right != nil {
		nr = n.Right.Size
	}
	n.Size = nl + nr + 1
	bt.updateSize(n.Parent)
}

func (bt *BinaryTree[T]) InsertAfter(old, new *node.Node[T]) error {
	if bt.Root == nil {
		return fmt.Errorf("empty tree")
	}
	if old == nil {
		return fmt.Errorf("empty node")
	}

	if old.Right == nil {
		old.Right = new
		new.Parent = old
		bt.updateSize(old)
		return nil
	}

	successor, err := bt.Successor(old)
	if err != nil {
		return err
	}
	successor.Left = new
	new.Parent = successor
	bt.updateSize(successor)
	return nil
}

func (bt *BinaryTree[T]) InsertBefore(old, new *node.Node[T]) error {
	if bt.Root == nil {
		return fmt.Errorf("empty tree")
	}
	if old == nil {
		return fmt.Errorf("empty node")
	}

	if old.Left == nil {
		old.Left = new
		new.Parent = old
		bt.updateSize(old)
		return nil
	}

	predecessor, err := bt.Predecessor(old)
	if err != nil {
		return err
	}
	predecessor.Right = new
	new.Parent = predecessor
	bt.updateSize(predecessor)
	return nil
}

func (bt *BinaryTree[T]) SubTree(n *node.Node[T], index int32) *node.Node[T] {
	var nl int32 = 0
	if n.Left != nil {
		nl = n.Left.Size
	}

	if index < nl {
		return bt.SubTree(n.Left, index)
	} else if index > nl {
		return bt.SubTree(n.Right, index-nl-1)
	}
	return n
}

func (bt *BinaryTree[T]) SubTreeFirst(n *node.Node[T]) (*node.Node[T], error) {
	if bt.GetSize() < 1 {
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
	if bt.GetSize() < 1 {
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
	if bt.GetSize() < 1 {
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
	if bt.GetSize() < 1 {
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
	if bt.GetSize() < 1 {
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
	if bt.GetSize() < 1 {
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
	if bt.GetSize() < 1 {
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
	if bt.Root == nil {
		return 0
	}
	return bt.Root.Size
}

func (bt *BinaryTree[T]) GetHeight() int32 {
	return bt.Height
}

// UTILITIES
func (bt *BinaryTree[T]) getNode(index int32) (*node.Node[T], error) {
	if index > bt.GetSize() {
		return nil, fmt.Errorf("index %d is larger than size %d", index, bt.GetSize())
	}

	return bt.SubTree(bt.Root, index), nil
}
