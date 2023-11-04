package binarysearchtree

import (
	"fmt"
	"math"

	"github.com/OladapoAjala/datastructures/queues/queue"
	"github.com/OladapoAjala/datastructures/sets"
	"github.com/OladapoAjala/datastructures/trees"
	"github.com/OladapoAjala/datastructures/trees/node"
	"golang.org/x/exp/constraints"
)

type BinarySearchTree[T constraints.Ordered] struct {
	Root   *node.Node[T]
	Size   int32
	Height int32
}

type IBinarySearchTree[T constraints.Ordered] interface {
	trees.ITrees[T]
	sets.Seter[T, any]
	// Insert(T) (*node.Node[T], error)
	// Delete(T) error
	// Find(T) (*node.Node[T], error)
}

// var _ IBinarySearchTree[string] = new(BinarySearchTree[string])

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return new(BinarySearchTree[T])
}

func (bst *BinarySearchTree[T]) InsertMany(data ...T) error {
	for _, d := range data {
		n, _ := bst.Find(d)
		if n != nil {
			return fmt.Errorf("%v present in tree", d)
		}
		_, err := bst.Insert(d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bst *BinarySearchTree[T]) Insert(data T) (*node.Node[T], error) {
	if data == *new(T) {
		return nil, fmt.Errorf("empty data")
	}

	if bst.Root == nil {
		v := node.NewNode[T](data)
		bst.Root = v
		bst.Size++
		return v, nil
	}
	v, err := bst.insert(data, bst.Root)
	if err != nil {
		return nil, err
	}
	bst.Size++
	return v, nil
}

func (bst *BinarySearchTree[T]) insert(data T, p *node.Node[T]) (*node.Node[T], error) {
	err := bst.validateData(data, p)
	if err != nil {
		return nil, err
	}

	if data < p.Data {
		if p.Left == nil {
			p.Left = node.NewNode(data)
			p.Left.Parent = p
			return p.Left, nil
		}
		return bst.insert(data, p.Left)
	}

	if p.Right == nil {
		p.Right = node.NewNode(data)
		p.Right.Parent = p
		return p.Right, nil
	}
	return bst.insert(data, p.Right)
}

func (bst *BinarySearchTree[T]) validateData(data T, p *node.Node[T]) error {
	if p == nil {
		return fmt.Errorf("empty node")
	}
	if data == p.Data {
		return fmt.Errorf("data %v already in tree", data)
	}
	return nil
}

func (bst *BinarySearchTree[T]) Find(data T) (*node.Node[T], error) {
	if data == *new(T) {
		return nil, fmt.Errorf("empty value")
	}
	if bst.Size == 0 {
		return nil, fmt.Errorf("empty tree")
	}
	return bst.find(data, bst.Root)
}

func (bst *BinarySearchTree[T]) find(data T, n *node.Node[T]) (*node.Node[T], error) {
	if n == nil {
		return nil, fmt.Errorf("data %v is not in tree", data)
	}

	if data < n.Data {
		return bst.find(data, n.Left)
	} else if data > n.Data {
		return bst.find(data, n.Right)
	}
	return n, nil
}

func (bst *BinarySearchTree[T]) Delete(data T) error {
	n, err := bst.Find(data)
	if err != nil {
		return err
	}
	return bst.delete(n)
}

func (bst *BinarySearchTree[T]) delete(n *node.Node[T]) error {
	if n.IsLeaf() {
		if n.Parent.Left == n {
			n.Parent.Left = nil
		} else {
			n.Parent.Right = nil
		}
		return nil
	}

	if n.Left != nil {
		pre, err := bst.Predecessor(n)
		if err != nil {
			return err
		}
		tmp := pre.Data
		pre.Data = n.Data
		n.Data = tmp
		return bst.delete(pre)
	}

	suc, err := bst.Successor(n)
	if err != nil {
		return err
	}
	tmp := suc.Data
	suc.Data = n.Data
	n.Data = tmp
	return bst.delete(suc)
}

func (bst *BinarySearchTree[T]) SubTreeFirst(n *node.Node[T]) (*node.Node[T], error) {
	if bst.Size < 1 {
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

func (bst *BinarySearchTree[T]) SubTreeLast(n *node.Node[T]) (*node.Node[T], error) {
	if bst.Size < 1 {
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

func (bst *BinarySearchTree[T]) Successor(n *node.Node[T]) (*node.Node[T], error) {
	if bst.Size < 1 {
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

func (bst *BinarySearchTree[T]) climbLeft(n *node.Node[T]) (*node.Node[T], error) {
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

func (bst *BinarySearchTree[T]) Predecessor(n *node.Node[T]) (*node.Node[T], error) {
	if bst.Size < 1 {
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

func (bst *BinarySearchTree[T]) climbRight(n *node.Node[T]) (*node.Node[T], error) {
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

func (bst *BinarySearchTree[T]) TraversalOrder(n *node.Node[T]) ([]T, error) {
	if bst.Size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []T{}, nil
	}

	leftOrder, err := bst.TraversalOrder(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := bst.TraversalOrder(n.Right)
	if err != nil {
		return nil, err
	}

	output := append(leftOrder, n.Data)
	output = append(output, rightOrder...)
	return output, nil
}

func (bst *BinarySearchTree[T]) PreOrderTraversal(n *node.Node[T]) ([]T, error) {
	if bst.Size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []T{}, nil
	}

	leftOrder, err := bst.PreOrderTraversal(n.Left)
	if err != nil {
		return nil, err
	}
	rightOrder, err := bst.PreOrderTraversal(n.Right)
	if err != nil {
		return nil, err
	}

	output := []T{n.Data}
	output = append(output, leftOrder...)
	output = append(output, rightOrder...)
	return output, nil
}

func (bst *BinarySearchTree[T]) PostOrderTraversal(n *node.Node[T]) ([]T, error) {
	if bst.Size < 1 {
		return nil, fmt.Errorf("empty tree")
	}
	if n == nil {
		return []T{}, nil
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
	output = append(output, n.Data)
	return output, nil
}

func (bst *BinarySearchTree[T]) GetSize() int32 {
	return bst.Size
}

func (bst *BinarySearchTree[T]) GetHeight() int32 {
	return height(bst.Root)
}

func height[T constraints.Ordered](node *node.Node[T]) int32 {
	if node == nil {
		return 0
	}

	height := math.Max(float64(height(node.GetLeft())), float64(height(node.GetRight()))) + 1
	return int32(height)
}

func findMin[T constraints.Ordered](node *node.Node[T]) *node.Node[T] {
	for node.GetLeft() != nil {
		node = node.GetLeft()
	}
	return node
}

func levelOrderTraversal[T constraints.Ordered](root *node.Node[T]) {
	if root == nil {
		return
	}

	height := height(root)
	var i int32
	for i = 0; i < height; i++ {
		fmt.Printf("Level %d: ", i)
		printLevel(i, root)
		fmt.Println()
	}
}

func printLevel[T constraints.Ordered](lvl int32, root *node.Node[T]) {
	if root == nil {
		return
	}

	if lvl == 0 {
		fmt.Printf("%v -> ", root.GetData())
	} else {
		lvl--
		printLevel(lvl, root.GetLeft())
		printLevel(lvl, root.GetRight())
	}
}

func breadthFirstSearch[T constraints.Ordered](bst *BinarySearchTree[T]) {
	que := queue.NewQueue[*node.Node[T]]()
	que.Enqueue(bst.Root)

	for {
		_node, err := que.Dequeue()
		if err != nil {
			return
		}

		fmt.Println(_node.GetData())

		if _node.GetLeft() != nil {
			que.Enqueue(_node.GetLeft())
		}

		if _node.GetRight() != nil {
			que.Enqueue(_node.GetRight())
		}
	}
}
