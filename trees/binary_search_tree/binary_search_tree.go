package binarysearchtree

import (
	"fmt"
	"math"

	"github.com/OladapoAjala/datastructures/queues/queue"
	"golang.org/x/exp/constraints"
)

const (
	preOrder = iota
	inOrder
	postOrder
	levelOrder
)

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
	Add(...T) error
	Contains(T) bool
	Remove(T) error
	Size() int32
	Height() int32
	Traverse(int)
}

var _ IBinarySearchTree[string] = new(BinarySearchTree[string])

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return new(BinarySearchTree[T])
}

func (bst *BinarySearchTree[T]) Add(data ...T) error {
	for _, d := range data {
		if isPresent := bst.Contains(d); isPresent {
			return fmt.Errorf("%v present in tree", d)
		}

		bst.Root = add(bst.Root, d)
		bst.nodeCount++
	}
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

func (bst *BinarySearchTree[T]) Remove(data T) error {
	if !bst.Contains(data) {
		return fmt.Errorf("%v not present in tree", data)
	}

	bst.Root = remove(bst.Root, data)
	bst.nodeCount--
	return nil
}

func remove[T constraints.Ordered](node *node[T], data T) *node[T] {
	if node == nil {
		return nil
	}

	if data < node.data {
		node.left = remove(node.left, data)
	} else if data > node.data {
		node.right = remove(node.right, data)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			tmp := findMin(node.right)
			node.data = tmp.data
			node.right = remove(node.right, tmp.data)
		}
	}

	return node
}

func findMin[T constraints.Ordered](node *node[T]) *node[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (bst *BinarySearchTree[T]) Size() int32 {
	return bst.nodeCount
}

func (bst *BinarySearchTree[T]) Height() int32 {
	return height(bst.Root)
}

func height[T constraints.Ordered](node *node[T]) int32 {
	if node == nil {
		return 0
	}

	height := math.Max(float64(height(node.left)), float64(height(node.right))) + 1
	return int32(height)
}

func (bst *BinarySearchTree[T]) Traverse(traversal int) {
	switch traversal {
	case preOrder:
		preOrderTraversal(bst.Root)
	case inOrder:
		inOrderTraversal(bst.Root)
	case postOrder:
		postOrderTraversal(bst.Root)
	case levelOrder:
		levelOrderTraversal(bst.Root)
	default:
		return
	}
}

func preOrderTraversal[T constraints.Ordered](root *node[T]) {
	if root == nil {
		return
	}

	fmt.Println(root.data)
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)
}

func inOrderTraversal[T constraints.Ordered](root *node[T]) {
	if root == nil {
		return
	}

	inOrderTraversal(root.left)
	fmt.Println(root.data)
	inOrderTraversal(root.right)
}

func postOrderTraversal[T constraints.Ordered](root *node[T]) {
	if root == nil {
		return
	}

	fmt.Println(root.data)
	postOrderTraversal(root.left)
	postOrderTraversal(root.right)
}

func levelOrderTraversal[T constraints.Ordered](root *node[T]) {
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

func printLevel[T constraints.Ordered](lvl int32, root *node[T]) {
	if root == nil {
		return
	}

	if lvl == 0 {
		fmt.Printf("%v -> ", root.data)
	} else {
		lvl--
		printLevel(lvl, root.left)
		printLevel(lvl, root.right)
	}
}

func breadthFirstSearch[T constraints.Ordered](bst *BinarySearchTree[T]) {
	que := queue.NewQueue[*node[T]]()
	que.Enqueue(bst.Root)

	for {
		_node, err := que.Dequeue()
		if err != nil {
			return
		}

		fmt.Println(_node.data)

		if _node.left != nil {
			que.Enqueue(_node.left)
		}

		if _node.right != nil {
			que.Enqueue(_node.right)
		}
	}
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

	3. Unused method
	func compare[T constraints.Ordered](a, b T) int32 {
		if a >= b {
			return GREATER
		}
		return LESSER
	}
***/
