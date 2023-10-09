package binarysearchtree

import (
	"fmt"
	"math"

	"github.com/OladapoAjala/datastructures/queues/queue"
	"github.com/OladapoAjala/datastructures/trees/node"
	"golang.org/x/exp/constraints"
)

const (
	preOrder = iota
	inOrder
	postOrder
	levelOrder
)

type BinarySearchTree[T constraints.Ordered] struct {
	Root      *node.Node[T]
	nodeCount int32
	height    int32
}

type IBinarySearchTree[T constraints.Ordered] interface {
	// trees.ITrees[T]
	Add(...T) error
	Contains(T) bool
	Remove(T) error
	Size() int32
	GetHeight() int32
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

func add[T constraints.Ordered](item *node.Node[T], data T) *node.Node[T] {
	if item == nil {
		return node.NewNode[T](data)
	}

	if data <= item.GetData() {
		item.Left = add(item.GetLeft(), data)
	} else {
		item.Right = add(item.GetRight(), data)
	}

	return item
}

func (bst *BinarySearchTree[T]) Contains(data T) bool {
	return contains(bst.Root, data)
}

func contains[T constraints.Ordered](item *node.Node[T], data T) bool {
	if item == nil {
		return false
	}

	if data < item.GetData() {
		return contains(item.GetLeft(), data)
	} else if data > item.GetData() {
		return contains(item.GetRight(), data)
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

func remove[T constraints.Ordered](item *node.Node[T], data T) *node.Node[T] {
	if item == nil {
		return nil
	}

	if data < item.GetData() {
		item.Left = remove(item.GetLeft(), data)
	} else if data > item.GetData() {
		item.Right = remove(item.GetRight(), data)
	} else {
		if item.GetLeft() == nil {
			return item.GetRight()
		} else if item.GetRight() == nil {
			return item.GetLeft()
		} else {
			tmp := findMin(item.GetRight())
			item.Data = tmp.GetData()
			item.Right = remove(item.GetRight(), tmp.GetData())
		}
	}

	return item
}

func findMin[T constraints.Ordered](node *node.Node[T]) *node.Node[T] {
	for node.GetLeft() != nil {
		node = node.GetLeft()
	}
	return node
}

func (bst *BinarySearchTree[T]) Size() int32 {
	return bst.nodeCount
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

func preOrderTraversal[T constraints.Ordered](root *node.Node[T]) {
	if root == nil {
		return
	}

	fmt.Println(root.GetData())
	preOrderTraversal(root.GetLeft())
	preOrderTraversal(root.GetRight())
}

func inOrderTraversal[T constraints.Ordered](root *node.Node[T]) {
	if root == nil {
		return
	}

	inOrderTraversal(root.GetLeft())
	fmt.Println(root.GetData())
	inOrderTraversal(root.GetRight())
}

func postOrderTraversal[T constraints.Ordered](root *node.Node[T]) {
	if root == nil {
		return
	}

	fmt.Println(root.GetData())
	postOrderTraversal(root.GetLeft())
	postOrderTraversal(root.GetRight())
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
