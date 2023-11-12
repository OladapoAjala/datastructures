package avltree

import (
	"fmt"
	"math"

	"github.com/OladapoAjala/datastructures/sets"
	"github.com/OladapoAjala/datastructures/trees"
	"github.com/OladapoAjala/datastructures/trees/data"
	"golang.org/x/exp/constraints"
)

type AVLTree[K constraints.Ordered, V comparable] struct {
	Root *data.Data[K, V]
}

type IAVLTree[K constraints.Ordered, V comparable] interface {
	trees.ITree[K, V]
	sets.Seter[K, V]
	InsertAfter(*data.Data[K, V], *data.Data[K, V]) error
	InsertBefore(*data.Data[K, V], *data.Data[K, V]) error
	SubTree(*data.Data[K, V], int32) *data.Data[K, V]
}

var _ IAVLTree[string, string] = new(AVLTree[string, string])

func NewAVLTree[K constraints.Ordered, V comparable](key K, val V) (*AVLTree[K, V], error) {
	avl := new(AVLTree[K, V])
	err := avl.Insert(key, val)
	if err != nil {
		return nil, err
	}
	return avl, nil
}

// SET METHODS
func (avl *AVLTree[K, V]) Insert(key K, val V) error {
	if key == *new(K) {
		return fmt.Errorf("empty key")
	}

	d := data.NewData[K, V](key, val)
	if avl.Root == nil {
		avl.Root = d
		return nil
	}
	_, err := avl.insert(d, avl.Root)
	if err != nil {
		return err
	}
	return nil
}

func (avl *AVLTree[K, V]) insert(new, root *data.Data[K, V]) (*data.Data[K, V], error) {
	err := avl.validateData(new, root)
	if err != nil {
		return nil, err
	}

	if new.GetKey() < root.GetKey() {
		if root.Left == nil {
			root.Left = new
			root.Left.Parent = root
			avl.maintain(new)
			return root.Left, nil
		}
		return avl.insert(new, root.Left)
	}

	if root.Right == nil {
		root.Right = new
		root.Right.Parent = root
		avl.maintain(new)
		return root.Right, nil
	}
	return avl.insert(new, root.Right)
}

func (avl *AVLTree[K, V]) validateData(new, root *data.Data[K, V]) error {
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

func (avl *AVLTree[K, V]) Find(key K) (V, error) {
	if key == *new(K) {
		return *new(V), fmt.Errorf("empty key")
	}
	if avl.GetSize() == 0 {
		return *new(V), fmt.Errorf("empty tree")
	}

	found, err := avl.find(key, avl.Root)
	if err != nil {
		return *new(V), err
	}
	return found.GetValue(), nil
}

func (avl *AVLTree[K, V]) find(key K, n *data.Data[K, V]) (*data.Data[K, V], error) {
	if n == nil {
		return nil, fmt.Errorf("key %v is not in tree", key)
	}
	if key < n.GetKey() {
		return avl.find(key, n.Left)
	} else if key > n.GetKey() {
		return avl.find(key, n.Right)
	}
	return n, nil
}

func (avl *AVLTree[K, V]) FindMax() (V, error) {
	max, err := avl.SubTreeLast(avl.Root)
	if err != nil {
		return *new(V), err
	}
	return max.GetValue(), nil
}

func (avl *AVLTree[K, V]) FindMin() (V, error) {
	min, err := avl.SubTreeFirst(avl.Root)
	if err != nil {
		return *new(V), err
	}
	return min.GetValue(), nil
}

func (avl *AVLTree[K, V]) FindNext(key K) (V, error) {
	if key == *new(K) {
		return *new(V), fmt.Errorf("empty key")
	}
	if avl.GetSize() == 0 {
		return *new(V), fmt.Errorf("empty tree")
	}

	curr, err := avl.find(key, avl.Root)
	if err != nil {
		return *new(V), err
	}
	successor, err := avl.Successor(curr)
	if err != nil {
		return *new(V), err
	}
	return successor.GetValue(), nil
}

func (avl *AVLTree[K, V]) FindPrev(key K) (V, error) {
	if key == *new(K) {
		return *new(V), fmt.Errorf("empty key")
	}
	if avl.GetSize() == 0 {
		return *new(V), fmt.Errorf("empty tree")
	}

	curr, err := avl.find(key, avl.Root)
	if err != nil {
		return *new(V), err
	}
	predecessor, err := avl.Predecessor(curr)
	if err != nil {
		return *new(V), err
	}
	return predecessor.GetValue(), nil
}

func (avl *AVLTree[K, V]) Delete(key K) (V, error) {
	n, err := avl.find(key, avl.Root)
	if err != nil {
		return *new(V), err
	}
	err = avl.delete(n)
	if err != nil {
		return *new(V), err
	}
	return n.GetValue(), nil
}

func (avl *AVLTree[K, V]) delete(n *data.Data[K, V]) error {
	if n.IsLeaf() {
		if n.Parent.Left == n {
			n.Parent.Left = nil
		} else {
			n.Parent.Right = nil
		}
		avl.maintain(n.Parent)
		return nil
	}

	if n.Left != nil {
		pre, err := avl.Predecessor(n)
		if err != nil {
			return err
		}
		tmp := pre.GetValue()
		pre.Value = n.GetValue()
		n.Value = tmp
		return avl.delete(pre)
	}

	suc, err := avl.Successor(n)
	if err != nil {
		return err
	}
	tmp := suc.GetValue()
	suc.Value = n.GetValue()
	n.Value = tmp
	return avl.delete(suc)
}

// TREE METHODS
func (avl *AVLTree[K, V]) update(n *data.Data[K, V]) {
	if n == nil {
		return
	}

	var sl, sr int32
	var hl, hr int32 = -1, -1
	if n.Left != nil {
		sl = n.Left.Size
		hl = n.Left.Height
	}
	if n.Right != nil {
		sr = n.Right.Size
		hr = n.Right.Height
	}
	n.Size = sl + sr + 1
	n.Height = 1 + int32(math.Max(float64(hl), float64(hr)))
}

func (avl *AVLTree[K, V]) balance(n *data.Data[K, V]) {
	if n == nil {
		return
	}

	skew := n.Skew()
	if skew == 2 {
		if n.Right.Skew() < 0 {
			avl.RotateRight(n.Right)
		}
		avl.RotateLeft(n)
	} else if skew == -2 {
		if n.Left.Skew() > 0 {
			avl.RotateLeft(n.Left)
		}
		avl.RotateRight(n)
	}
}

func (avl *AVLTree[K, V]) maintain(n *data.Data[K, V]) {
	if n == nil {
		return
	}
	avl.balance(n)
	avl.update(n)
	avl.maintain(n.Parent)
}

func (avl *AVLTree[K, V]) RotateRight(n *data.Data[K, V]) error {
	if n.Left == nil {
		return fmt.Errorf("node %v has no left child", n.GetValue())
	}

	parent := n.Parent
	left := n.Left
	n.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = n
	}
	left.Right = n
	n.Parent = left
	left.Parent = parent
	if parent == nil {
		avl.Root = left
		return nil
	}

	if parent.Left == n {
		parent.Left = left
	} else {
		parent.Right = left
	}
	return nil
}

func (avl *AVLTree[K, V]) RotateLeft(n *data.Data[K, V]) error {
	if n.Right == nil {
		return fmt.Errorf("node %v has no right child", n.GetValue())
	}

	parent := n.Parent
	right := n.Right
	n.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = n
	}
	right.Left = n
	n.Parent = right
	right.Parent = parent
	if parent == nil {
		avl.Root = right
		return nil
	}

	if parent.Right == n {
		parent.Right = right
	} else {
		parent.Left = right
	}
	return nil
}

func (avl *AVLTree[K, V]) InsertAfter(old, new *data.Data[K, V]) error {
	if avl.Root == nil {
		return fmt.Errorf("empty tree")
	}
	if old == nil {
		return fmt.Errorf("empty node")
	}

	if old.Right == nil {
		old.Right = new
		new.Parent = old
		avl.maintain(old)
		return nil
	}

	successor, err := avl.Successor(old)
	if err != nil {
		return err
	}
	successor.Left = new
	new.Parent = successor
	avl.maintain(successor)
	return nil
}

func (avl *AVLTree[K, V]) InsertBefore(old, new *data.Data[K, V]) error {
	if avl.Root == nil {
		return fmt.Errorf("empty tree")
	}
	if old == nil {
		return fmt.Errorf("empty node")
	}

	if old.Left == nil {
		old.Left = new
		new.Parent = old
		avl.maintain(old)
		return nil
	}

	predecessor, err := avl.Predecessor(old)
	if err != nil {
		return err
	}
	predecessor.Right = new
	new.Parent = predecessor
	avl.maintain(predecessor)
	return nil
}

func (avl *AVLTree[K, V]) SubTree(n *data.Data[K, V], index int32) *data.Data[K, V] {
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

func (avl *AVLTree[K, V]) SubTreeFirst(n *data.Data[K, V]) (*data.Data[K, V], error) {
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

func (avl *AVLTree[K, V]) SubTreeLast(n *data.Data[K, V]) (*data.Data[K, V], error) {
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

func (avl *AVLTree[K, V]) Successor(n *data.Data[K, V]) (*data.Data[K, V], error) {
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

func (avl *AVLTree[K, V]) climbLeft(n *data.Data[K, V]) (*data.Data[K, V], error) {
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

func (avl *AVLTree[K, V]) Predecessor(n *data.Data[K, V]) (*data.Data[K, V], error) {
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

func (avl *AVLTree[K, V]) climbRight(n *data.Data[K, V]) (*data.Data[K, V], error) {
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

func (avl *AVLTree[K, V]) TraversalOrder(n *data.Data[K, V]) ([]V, error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []V{}, nil
	}

	leftOrder, err := avl.TraversalOrder(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := avl.TraversalOrder(n.Right)
	if err != nil {
		return nil, err
	}

	output := append(leftOrder, n.GetValue())
	output = append(output, rightOrder...)
	return output, nil
}

func (avl *AVLTree[K, V]) PreOrderTraversal(n *data.Data[K, V]) ([]V, error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []V{}, nil
	}

	leftOrder, err := avl.PreOrderTraversal(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := avl.PreOrderTraversal(n.Right)
	if err != nil {
		return nil, err
	}

	output := []V{n.GetValue()}
	output = append(output, leftOrder...)
	output = append(output, rightOrder...)
	return output, nil
}

func (avl *AVLTree[K, V]) PostOrderTraversal(n *data.Data[K, V]) ([]V, error) {
	if avl.GetSize() < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []V{}, nil
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
	output = append(output, n.GetValue())
	return output, nil
}

func (avl *AVLTree[K, V]) Size() int32 {
	return avl.GetSize()
}

func (avl *AVLTree[K, V]) GetSize() int32 {
	if avl.Root == nil {
		return 0
	}
	return avl.Root.Size
}

func (avl *AVLTree[K, V]) IsEmpty() bool {
	return avl.Root == nil
}

func (avl *AVLTree[K, V]) GetHeight() int32 {
	return avl.Root.Height
}
