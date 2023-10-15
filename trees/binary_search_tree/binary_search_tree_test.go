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

func Test_Contains(t *testing.T) {
	is := assert.New(t)
	bst := NewBinarySearchTree[string]()

	type args struct {
		data string
	}

	tests := []struct {
		name  string
		args  args
		bst   *BinarySearchTree[string]
		setup func(*BinarySearchTree[string])
		want  func(bool)
	}{
		{
			name: "search empty BST",
			args: args{
				data: "A",
			},
			bst: bst,
			want: func(b bool) {
				is.False(b)
			},
		},
		{
			name: "find A in BST",
			args: args{
				data: "A",
			},
			bst: bst,
			setup: func(bst *BinarySearchTree[string]) {
				bst.InsertMany("B")
				bst.InsertMany("A")
				bst.InsertMany("C")
			},
			want: func(b bool) {
				is.True(b)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.bst)
			}
			isPresent := tt.bst.Contains(tt.args.data)
			tt.want(isPresent)
		})
	}
}

func Test_Remove(t *testing.T) {
	is := assert.New(t)
	bst := NewBinarySearchTree[string]()

	type args struct {
		data string
	}

	tests := []struct {
		name  string
		args  args
		bst   *BinarySearchTree[string]
		setup func(*BinarySearchTree[string])
		want  func(*BinarySearchTree[string], error)
	}{
		{
			name: "simple removal",
			args: args{
				data: "A",
			},
			bst: bst,
			setup: func(bst *BinarySearchTree[string]) {
				bst.InsertMany("B")
				bst.InsertMany("A")
				bst.InsertMany("C")
			},
			want: func(bst *BinarySearchTree[string], err error) {
				is.Nil(err)
				is.Equal(bst.Root.GetData(), "B")
				is.Nil(bst.Root.GetLeft())
				is.Equal(bst.Root.GetRight().GetData(), "C")
				is.EqualValues(bst.GetSize(), 2)
			},
		},
		{
			name: "complex removal",
			args: args{
				data: "E",
			},
			bst: bst,
			setup: func(bst *BinarySearchTree[string]) {
				bst.InsertMany("E")
				bst.InsertMany("D")
				bst.InsertMany("I")
				bst.InsertMany("F")
				bst.InsertMany("J")
			},
			want: func(bst *BinarySearchTree[string], err error) {
				is.Nil(err)
				is.Equal(bst.Root.GetData(), "B")
				is.Nil(bst.Root.GetLeft())
				is.Equal(bst.Root.GetRight().GetData(), "C")
				is.Equal(bst.Root.GetRight().GetRight().GetData(), "F")
				is.Equal(bst.Root.GetRight().GetRight().GetLeft().GetData(), "D")
				is.Equal(bst.Root.GetRight().GetRight().GetRight().GetData(), "I")
				is.Nil(bst.Root.GetRight().GetRight().GetRight().GetLeft())
				is.Equal(bst.Root.GetRight().GetRight().GetRight().GetRight().GetData(), "J")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.bst)
			}
			err := tt.bst.Remove(tt.args.data)
			tt.want(tt.bst, err)
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
