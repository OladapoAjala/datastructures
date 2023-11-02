package avltree

import (
	"fmt"
	"math"

	"github.com/OladapoAjala/datastructures/sequences"
	"github.com/OladapoAjala/datastructures/trees"
	"github.com/OladapoAjala/datastructures/trees/node"
)

type AVLTree[T comparable] struct {
	Root *node.Node[T]
}

type IAVLTree[T comparable] interface {
	trees.ITrees[T]
	sequences.Sequencer[T]
	InsertAfter(*node.Node[T], *node.Node[T]) error
	InsertBefore(*node.Node[T], *node.Node[T]) error
	SubTree(*node.Node[T], int32) *node.Node[T]
}

var _ IAVLTree[string] = new(AVLTree[string])

// SEQUENCE METHODS
func (avl *AVLTree[T]) GetData(index int32) (T, error) {
	n, err := avl.getNode(index)
	if err != nil {
		return *new(T), err
	}
	return n.GetData(), nil
}

func (avl *AVLTree[T]) Contains(data T) bool {
	return avl.contains(avl.Root, data)
}

func (avl *AVLTree[T]) contains(n *node.Node[T], data T) bool {
	if n == nil {
		return false
	}
	if n.Data == data {
		return true
	}

	left := avl.contains(n.Left, data)
	if left {
		return left
	}
	return avl.contains(n.Right, data)
}

func (avl *AVLTree[T]) Insert(index int32, data T) error {
	if data == *new(T) {
		return fmt.Errorf("empty data")
	}
	if index > avl.GetSize() {
		return fmt.Errorf("index %d is larger than size %d", index, avl.GetSize())
	}

	newNode := node.NewNode[T](data)
	if avl.Root == nil {
		avl.Root = newNode
		return nil
	}
	if index == avl.GetSize() {
		last := avl.SubTree(avl.Root, avl.GetSize()-1)
		last.Right = newNode
		newNode.Parent = last
		avl.update(last)
		return nil
	}

	sub := avl.SubTree(avl.Root, index)
	err := avl.InsertBefore(sub, newNode)
	if err != nil {
		return err
	}
	return nil
}

func (avl *AVLTree[T]) Delete(index int32) error {
	if avl.Root == nil {
		return fmt.Errorf("cannot delete from empty tree")
	}

	n, err := avl.getNode(index)
	if err != nil {
		return err
	}
	return avl.delete(n)
}

func (avl *AVLTree[T]) delete(n *node.Node[T]) error {
	if n.IsLeaf() {
		if n == avl.Root {
			avl.Root = nil
		} else if n.Parent.Left == n {
			n.Parent.Left = nil
		} else {
			n.Parent.Right = nil
		}
		avl.update(n.Parent)
		return nil
	}

	if n.Left != nil {
		pre, err := avl.Predecessor(n)
		if err != nil {
			return err
		}
		tmp := pre.Data
		pre.Data = n.Data
		n.Data = tmp
		return avl.delete(pre)
	}

	suc, err := avl.Successor(n)
	if err != nil {
		return err
	}
	tmp := suc.Data
	suc.Data = n.Data
	n.Data = tmp
	return avl.delete(suc)
}

func (avl *AVLTree[T]) Set(index int32, data T) error {
	n, err := avl.getNode(index)
	if err != nil {
		return err
	}
	n.Data = data
	return nil
}

func (avl *AVLTree[T]) InsertFirst(data T) error {
	return avl.Insert(0, data)
}

func (avl *AVLTree[T]) InsertLast(data T) error {
	return avl.Insert(avl.GetSize(), data)
}

func (avl *AVLTree[T]) DeleteFirst() error {
	return avl.Delete(0)
}

func (avl *AVLTree[T]) DeleteLast() error {
	return avl.Delete(avl.GetSize() - 1)
}

// TREE METHODS
func NewAVLTree[T comparable](data ...T) *AVLTree[T] {
	avl := new(AVLTree[T])
	for i, item := range data {
		err := avl.Insert(int32(i), item)
		if err != nil {
			return nil
		}
	}
	return avl
}

func (avl *AVLTree[T]) update(n *node.Node[T]) {
	if n == nil {
		return
	}
	var nl, nr int32
	var hl, hr int32
	if n.Left != nil {
		nl = n.Left.Size
		hl = n.Left.Height
	} else {
		hl = -1
	}
	if n.Right != nil {
		nr = n.Right.Size
		hr = n.Right.Height
	} else {
		hr = -1
	}
	n.Size = nl + nr + 1
	n.Height = 1 + int32(math.Max(float64(hl), float64(hr)))
	avl.balance(n)
	avl.update(n.Parent)
}

func (avl *AVLTree[T]) balance(n *node.Node[T]) {
	skew := n.Skew()
	if skew == 2 {
		if n.Right.Skew() >= 0 {
			// right=right case
			n.LeftRotate()
		} else {
			// right-left case
			n.Right.RightRotate()
			n.LeftRotate()
		}
	} else if skew == -2 {
		if n.Left.Skew() <= 0 {
			// left-left case
			n.RightRotate()
		} else {
			// left-right case
			n.Left.LeftRotate()
			n.RightRotate()
		}
	}
}

func (avl *AVLTree[T]) InsertAfter(old, new *node.Node[T]) error {
	if avl.Root == nil {
		return fmt.Errorf("empty tree")
	}
	if old == nil {
		return fmt.Errorf("empty node")
	}

	if old.Right == nil {
		old.Right = new
		new.Parent = old
		avl.update(old)
		return nil
	}

	successor, err := avl.Successor(old)
	if err != nil {
		return err
	}
	successor.Left = new
	new.Parent = successor
	avl.update(successor)
	return nil
}

func (avl *AVLTree[T]) InsertBefore(old, new *node.Node[T]) error {
	if avl.Root == nil {
		return fmt.Errorf("empty tree")
	}
	if old == nil {
		return fmt.Errorf("empty node")
	}

	if old.Left == nil {
		old.Left = new
		new.Parent = old
		avl.update(old)
		return nil
	}

	predecessor, err := avl.Predecessor(old)
	if err != nil {
		return err
	}
	predecessor.Right = new
	new.Parent = predecessor
	avl.update(predecessor)
	return nil
}

func (avl *AVLTree[T]) SubTree(n *node.Node[T], index int32) *node.Node[T] {
	var nl int32 = 0
	if n.Left != nil {
		nl = n.Left.Size
	}

	if index < nl {
		return avl.SubTree(n.Left, index)
	} else if index > nl {
		return avl.SubTree(n.Right, index-nl-1)
	}
	return n
}

func (avl *AVLTree[T]) SubTreeFirst(n *node.Node[T]) (*node.Node[T], error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Left == nil {
		return n, nil
	}
	return avl.SubTreeFirst(n.Left)
}

func (avl *AVLTree[T]) SubTreeLast(n *node.Node[T]) (*node.Node[T], error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Right == nil {
		return n, nil
	}
	return avl.SubTreeLast(n.Right)
}

func (avl *AVLTree[T]) Successor(n *node.Node[T]) (*node.Node[T], error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Right == nil {
		return avl.climbLeft(n)
	}
	return avl.SubTreeFirst(n.Right)
}

func (avl *AVLTree[T]) climbLeft(n *node.Node[T]) (*node.Node[T], error) {
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}
	if n.Parent == nil {
		return nil, fmt.Errorf("node %v has no parent", n)
	}

	if n == n.Parent.Left {
		return n.Parent, nil
	}
	return avl.climbLeft(n.Parent)
}

func (avl *AVLTree[T]) Predecessor(n *node.Node[T]) (*node.Node[T], error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Left == nil {
		return avl.climbRight(n)
	}
	return avl.SubTreeLast(n.Left)
}

func (avl *AVLTree[T]) climbRight(n *node.Node[T]) (*node.Node[T], error) {
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}
	if n.Parent == nil {
		return nil, fmt.Errorf("node %v has no parent", n)
	}

	if n == n.Parent.Right {
		return n.Parent, nil
	}
	return avl.climbRight(n.Parent)
}

func (avl *AVLTree[T]) TraversalOrder(n *node.Node[T]) ([]T, error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []T{}, nil
	}

	leftOrder, err := avl.TraversalOrder(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := avl.TraversalOrder(n.Right)
	if err != nil {
		return nil, err
	}

	output := append(leftOrder, n.Data)
	output = append(output, rightOrder...)
	return output, nil
}

func (avl *AVLTree[T]) PreOrderTraversal(n *node.Node[T]) ([]T, error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []T{}, nil
	}

	leftOrder, err := avl.PreOrderTraversal(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := avl.PreOrderTraversal(n.Right)
	if err != nil {
		return nil, err
	}

	output := []T{n.Data}
	output = append(output, leftOrder...)
	output = append(output, rightOrder...)
	return output, nil
}

func (avl *AVLTree[T]) PostOrderTraversal(n *node.Node[T]) ([]T, error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []T{}, nil
	}

	leftOrder, err := avl.PostOrderTraversal(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := avl.PostOrderTraversal(n.Right)
	if err != nil {
		return nil, err
	}

	output := append(leftOrder, rightOrder...)
	output = append(output, n.Data)
	return output, nil
}

func (avl *AVLTree[T]) GetSize() int32 {
	if avl.Root == nil {
		return 0
	}
	return avl.Root.Size
}

func (avl *AVLTree[T]) IsEmpty() bool {
	return avl.Root == nil
}

func (avl *AVLTree[T]) GetHeight() int32 {
	return avl.Root.Height
}

// UTILITIES
func (avl *AVLTree[T]) getNode(index int32) (*node.Node[T], error) {
	if index >= avl.GetSize() {
		return nil, fmt.Errorf("index %d is out of range", index)
	}
	return avl.SubTree(avl.Root, index), nil
}
