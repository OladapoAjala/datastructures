package binarysearchtree

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/trees/node"
	"github.com/stretchr/testify/assert"
)

func Test_NewBinarySearchTree(t *testing.T) {
	is := assert.New(t)
	bst := NewBinarySearchTree[string]()

	is.EqualValues(bst.GetSize(), 0)
	is.Nil(bst.Root)
}

func Test_Insert(t *testing.T) {
	is := assert.New(t)
	bst := NewBinarySearchTree[int]()

	tests := []struct {
		name  string
		input int
		want  func(*node.Node[int], error)
	}{
		{
			name:  "Insert into an empty tree",
			input: 10,
			want: func(n *node.Node[int], err error) {
				is.Nil(err)
				is.Equal(bst.Root, n)
				is.Nil(bst.Root.Left)
				is.Nil(bst.Root.Right)
				is.EqualValues(1, bst.GetSize())
			},
		},
		{
			name:  "Insert smaller value to the left",
			input: 5,
			want: func(n *node.Node[int], err error) {
				is.Nil(err)
				is.Equal(bst.Root.Left, n)
				is.Nil(bst.Root.Right)
				is.EqualValues(2, bst.GetSize())
			},
		},
		{
			name:  "Insert larger value to the right",
			input: 15,
			want: func(n *node.Node[int], err error) {
				is.Nil(err)
				is.Equal(bst.Root.Left.Data, 5)
				is.Equal(bst.Root.Right, n)
				is.EqualValues(3, bst.GetSize())
			},
		},
		{
			name:  "Insert duplicate value",
			input: 10,
			want: func(n *node.Node[int], err error) {
				is.Nil(n)
				is.Error(fmt.Errorf("data 10 already in tree"))
				is.EqualValues(3, bst.GetSize())
			},
		},
		{
			name:  "Insert a value into left sub-tree",
			input: 3,
			want: func(n *node.Node[int], err error) {
				is.Nil(err)
				is.Equal(bst.Root.Left, n.Parent)
				is.True(n.IsLeaf())
				is.EqualValues(4, bst.GetSize())
			},
		},
		{
			name:  "Insert a value into left sub-tree",
			input: 12,
			want: func(n *node.Node[int], err error) {
				is.Nil(err)
				is.Equal(bst.Root.Right, n.Parent)
				is.True(n.IsLeaf())
				is.EqualValues(5, bst.GetSize())
			},
		},
		{
			name:  "Insert negative values",
			input: -10,
			want: func(n *node.Node[int], err error) {
				is.Nil(err)
				is.Equal(bst.Root.Left.Left, n.Parent)
				is.True(n.IsLeaf())
				is.EqualValues(6, bst.GetSize())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, err := bst.Insert(tt.input)
			tt.want(n, err)
		})
	}
}

func Test_InsertMany(t *testing.T) {
	is := assert.New(t)
	bst := NewBinarySearchTree[int]()

	tests := []struct {
		name string
		data []int
		bst  *BinarySearchTree[int]
		want func(*BinarySearchTree[int], error)
	}{
		{
			name: "insert an array of data",
			data: []int{3, 12, 8, 11, 1, 4, 2, 5},
			bst:  bst,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				is.EqualValues(bst.GetSize(), 8)

				is.Equal(bst.Root.GetData(), 3)
				is.Equal(bst.Root.Left.GetData(), 1)
				is.Equal(bst.Root.Left.Right.GetData(), 2)

				is.Equal(bst.Root.Right.GetData(), 12)
				is.Equal(bst.Root.Right.Left.GetData(), 8)
				is.Equal(bst.Root.Right.Left.Left.GetData(), 4)
				is.Equal(bst.Root.Right.Left.Right.GetData(), 11)
				is.Equal(bst.Root.Right.Left.Left.Right.GetData(), 5)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.bst.InsertMany(tt.data...)
			tt.want(tt.bst, err)
		})
	}
}

func Test_Delete(t *testing.T) {
	is := assert.New(t)
	bst := NewBinarySearchTree[int]()

	tests := []struct {
		name  string
		setup func(bst *BinarySearchTree[int])
		arg   int
		want  func(*BinarySearchTree[int], error)
	}{
		{
			name: "Delete leaf node",
			setup: func(bst *BinarySearchTree[int]) {
				err := bst.InsertMany(10, 5, 15, 3, 8, 12, 20)
				is.Nil(err)
			},
			arg: 3,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				order, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(order, []int{5, 8, 10, 12, 15, 20})
			},
		},
		{
			name: "Delete node with one child",
			arg:  5,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				order, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(order, []int{8, 10, 12, 15, 20})
			},
		},
		{
			name: "Delete node with two child",
			arg:  15,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				order, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(order, []int{8, 10, 12, 20})
			},
		},
		{
			name: "Delete root node",
			arg:  10,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				order, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(order, []int{8, 12, 20})
			},
		},
		{
			name: "Delete non-existent value",
			arg:  7,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Error(err, fmt.Errorf("data 7 is not in tree"))
				order, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(order, []int{8, 12, 20})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(bst)
			}
			err := bst.Delete(tt.arg)
			tt.want(bst, err)
		})
	}
}

func Test_Height(t *testing.T) {
	is := assert.New(t)
	bst := NewBinarySearchTree[string]()

	tests := []struct {
		name  string
		bst   *BinarySearchTree[string]
		setup func(*BinarySearchTree[string])
		want  func(int32)
	}{
		{
			name: "simple BST",
			bst:  bst,
			setup: func(bst *BinarySearchTree[string]) {
				bst.InsertMany("B")
				bst.InsertMany("A")
				bst.InsertMany("C")
			},
			want: func(height int32) {
				is.EqualValues(height, 2)
			},
		},
		{
			name: "longer right sub-tree",
			bst:  bst,
			setup: func(bst *BinarySearchTree[string]) {
				bst.InsertMany("D")
			},
			want: func(height int32) {
				is.EqualValues(height, 3)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.bst)
			}
			height := tt.bst.GetHeight()
			tt.want(height)
		})
	}
}
