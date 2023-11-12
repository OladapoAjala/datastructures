package binarysearchtree

import (
	"fmt"
	"math"

	"github.com/OladapoAjala/datastructures/sets"
	"github.com/OladapoAjala/datastructures/trees"
	"github.com/OladapoAjala/datastructures/trees/data"
	"golang.org/x/exp/constraints"
)

type BinarySearchTree[K constraints.Ordered, V comparable] struct {
	Root *data.Data[K, V]
}

type IBinarySearchTree[K constraints.Ordered, V comparable] interface {
	trees.ITree[K, V]
	sets.Seter[K, V]
}

var _ IBinarySearchTree[string, string] = new(BinarySearchTree[string, string])

func NewBSTTree[K constraints.Ordered, V comparable](key K, val V) (*BinarySearchTree[K, V], error) {
	bst := new(BinarySearchTree[K, V])
	err := bst.Insert(key, val)
	if err != nil {
		return nil, err
	}
	return bst, nil
}

// SET METHODS
func (bst *BinarySearchTree[K, V]) Insert(key K, val V) error {
	if key == *new(K) {
		return fmt.Errorf("empty key")
	}

	d := data.NewData[K, V](key, val)
	if bst.Root == nil {
		bst.Root = d
		return nil
	}
	_, err := bst.insert(d, bst.Root)
	if err != nil {
		return err
	}
	return nil
}

func (bst *BinarySearchTree[K, V]) insert(new, root *data.Data[K, V]) (*data.Data[K, V], error) {
	err := bst.validateData(new, root)
	if err != nil {
		return nil, err
	}

	if new.GetKey() < root.GetKey() {
		if root.Left == nil {
			root.Left = new
			root.Left.Parent = root
			bst.update(new)
			return root.Left, nil
		}
		return bst.insert(new, root.Left)
	}

	if root.Right == nil {
		root.Right = new
		root.Right.Parent = root
		bst.update(new)
		return root.Right, nil
	}
	return bst.insert(new, root.Right)
}

func (bst *BinarySearchTree[K, V]) validateData(new, root *data.Data[K, V]) error {
	if root == nil {
		return fmt.Errorf("empty node")
	}
	if new.IsEqual(root) {
		return fmt.Errorf("key %v already in tree", new)
	}
	if new.IsEmpty() {
		return fmt.Errorf("data %v is empty", new)
	}
	return nil
}

func (bst *BinarySearchTree[K, V]) Find(key K) (V, error) {
	if key == *new(K) {
		return *new(V), fmt.Errorf("empty key")
	}
	if bst.GetSize() == 0 {
		return *new(V), fmt.Errorf("empty tree")
	}

	found, err := bst.find(key, bst.Root)
	if err != nil {
		return *new(V), err
	}
	return found.GetValue(), nil
}

func (bst *BinarySearchTree[K, V]) find(key K, n *data.Data[K, V]) (*data.Data[K, V], error) {
	if n == nil {
		return nil, fmt.Errorf("key %v is not in tree", key)
	}
	if key < n.GetKey() {
		return bst.find(key, n.Left)
	} else if key > n.GetKey() {
		return bst.find(key, n.Right)
	}
	return n, nil
}

func (bst *BinarySearchTree[K, V]) FindMax() (V, error) {
	max, err := bst.SubTreeLast(bst.Root)
	if err != nil {
		return *new(V), err
	}
	return max.GetValue(), nil
}

func (bst *BinarySearchTree[K, V]) FindMin() (V, error) {
	min, err := bst.SubTreeFirst(bst.Root)
	if err != nil {
		return *new(V), err
	}
	return min.GetValue(), nil
}

func (bst *BinarySearchTree[K, V]) FindNext(key K) (V, error) {
	if key == *new(K) {
		return *new(V), fmt.Errorf("empty key")
	}
	if bst.GetSize() == 0 {
		return *new(V), fmt.Errorf("empty tree")
	}

	curr, err := bst.find(key, bst.Root)
	if err != nil {
		return *new(V), err
	}
	successor, err := bst.Successor(curr)
	if err != nil {
		return *new(V), err
	}
	return successor.GetValue(), nil
}

func (bst *BinarySearchTree[K, V]) FindPrev(key K) (V, error) {
	if key == *new(K) {
		return *new(V), fmt.Errorf("empty key")
	}
	if bst.GetSize() == 0 {
		return *new(V), fmt.Errorf("empty tree")
	}

	curr, err := bst.find(key, bst.Root)
	if err != nil {
		return *new(V), err
	}
	predecessor, err := bst.Predecessor(curr)
	if err != nil {
		return *new(V), err
	}
	return predecessor.GetValue(), nil
}

func (bst *BinarySearchTree[K, V]) Delete(key K) (V, error) {
	n, err := bst.find(key, bst.Root)
	if err != nil {
		return *new(V), err
	}
	err = bst.delete(n)
	if err != nil {
		return *new(V), err
	}
	return n.GetValue(), nil
}

func (bst *BinarySearchTree[K, V]) delete(n *data.Data[K, V]) error {
	if n.IsLeaf() {
		if n.Parent.Left == n {
			n.Parent.Left = nil
		} else {
			n.Parent.Right = nil
		}
		bst.update(n.Parent)
		return nil
	}

	if n.Left != nil {
		pre, err := bst.Predecessor(n)
		if err != nil {
			return err
		}
		tmp := pre.GetValue()
		pre.Value = n.GetValue()
		n.Value = tmp
		return bst.delete(pre)
	}

	suc, err := bst.Successor(n)
	if err != nil {
		return err
	}
	tmp := suc.GetValue()
	suc.Value = n.GetValue()
	n.Value = tmp
	return bst.delete(suc)
}

// TREE METHODS
func (bst *BinarySearchTree[K, V]) update(n *data.Data[K, V]) {
	if n == nil {
		return
	}
	var sl, sr int32
	var hl, hr int32
	if n.Left != nil {
		sl = n.Left.Size
		hl = n.Left.Height
	} else {
		hl = -1
	}
	if n.Right != nil {
		sr = n.Right.Size
		hr = n.Right.Height
	} else {
		hr = -1
	}
	n.Size = sl + sr + 1
	n.Height = 1 + int32(math.Max(float64(hl), float64(hr)))
	bst.update(n.Parent)
}

func (bst *BinarySearchTree[K, V]) SubTreeFirst(n *data.Data[K, V]) (*data.Data[K, V], error) {
	if bst.IsEmpty() {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Left == nil {
		return n, nil
	}
	return bst.SubTreeFirst(n.Left)
}

func (bst *BinarySearchTree[K, V]) SubTreeLast(n *data.Data[K, V]) (*data.Data[K, V], error) {
	if bst.IsEmpty() {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Right == nil {
		return n, nil
	}
	return bst.SubTreeLast(n.Right)
}

func (bst *BinarySearchTree[K, V]) Successor(n *data.Data[K, V]) (*data.Data[K, V], error) {
	if bst.IsEmpty() {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Right == nil {
		return bst.climbLeft(n)
	}
	return bst.SubTreeFirst(n.Right)
}

func (bst *BinarySearchTree[K, V]) climbLeft(n *data.Data[K, V]) (*data.Data[K, V], error) {
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}
	if n.Parent == nil {
		return nil, fmt.Errorf("node %v has no parent", n)
	}

	if n == n.Parent.Left {
		return n.Parent, nil
	}
	return bst.climbLeft(n.Parent)
}

func (bst *BinarySearchTree[K, V]) Predecessor(n *data.Data[K, V]) (*data.Data[K, V], error) {
	if bst.IsEmpty() {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}

	if n.Left == nil {
		return bst.climbRight(n)
	}
	return bst.SubTreeLast(n.Left)
}

func (bst *BinarySearchTree[K, V]) climbRight(n *data.Data[K, V]) (*data.Data[K, V], error) {
	if n == nil {
		return nil, fmt.Errorf("empty node")
	}
	if n.Parent == nil {
		return nil, fmt.Errorf("node %v has no parent", n)
	}

	if n == n.Parent.Right {
		return n.Parent, nil
	}
	return bst.climbRight(n.Parent)
}

func (bst *BinarySearchTree[K, V]) TraversalOrder(n *data.Data[K, V]) ([]V, error) {
	if bst.IsEmpty() {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []V{}, nil
	}

	leftOrder, err := bst.TraversalOrder(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := bst.TraversalOrder(n.Right)
	if err != nil {
		return nil, err
	}

	output := append(leftOrder, n.GetValue())
	output = append(output, rightOrder...)
	return output, nil
}

func (bst *BinarySearchTree[K, V]) PreOrderTraversal(n *data.Data[K, V]) ([]V, error) {
	if bst.IsEmpty() {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []V{}, nil
	}

	leftOrder, err := bst.PreOrderTraversal(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := bst.PreOrderTraversal(n.Right)
	if err != nil {
		return nil, err
	}

	output := []V{n.GetValue()}
	output = append(output, leftOrder...)
	output = append(output, rightOrder...)
	return output, nil
}

func (bst *BinarySearchTree[K, V]) PostOrderTraversal(n *data.Data[K, V]) ([]V, error) {
	if bst.IsEmpty() {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []V{}, nil
	}

	leftOrder, err := bst.PostOrderTraversal(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := bst.PostOrderTraversal(n.Right)
	if err != nil {
		return nil, err
	}

	output := append(leftOrder, rightOrder...)
	output = append(output, n.GetValue())
	return output, nil
}

func (bst *BinarySearchTree[K, V]) Size() int32 {
	return bst.GetSize()
}

func (bst *BinarySearchTree[K, V]) GetSize() int32 {
	if bst.Root == nil {
		return 0
	}
	return bst.Root.Size
}

func (bst *BinarySearchTree[K, V]) IsEmpty() bool {
	return bst.Root == nil
}

func (bst *BinarySearchTree[K, V]) GetHeight() int32 {
	return bst.Root.Height
}
