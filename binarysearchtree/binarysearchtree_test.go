package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewBinarySearchTree(t *testing.T) {
	is := assert.New(t)
	bst := NewBinarySearchTree[string]()

	is.EqualValues(bst.nodeCount, 0)
	is.Nil(bst.Root)
}

func Test_Add(t *testing.T) {
	is := assert.New(t)
	bst := NewBinarySearchTree[int]()

	type args struct {
		data int
	}

	tests := []struct {
		name string
		args args
		bst  *BinarySearchTree[int]
		want func(*BinarySearchTree[int], error)
	}{
		{
			name: "insert root element",
			args: args{
				data: 4,
			},
			bst: bst,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				is.Equal(bst.Root.data, 4)
				is.EqualValues(bst.nodeCount, 1)
			},
		},
		{
			name: "insert element lesser than root",
			args: args{
				data: 2,
			},
			bst: bst,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				is.EqualValues(bst.nodeCount, 2)
				is.NotNil(bst.Root.left)
				is.Nil(bst.Root.right)
				is.Equal(bst.Root.left.data, 2)
			},
		},
		{
			name: "insert element greater than root",
			args: args{
				data: 5,
			},
			bst: bst,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				is.EqualValues(bst.nodeCount, 3)
				is.NotNil(bst.Root.left)
				is.NotNil(bst.Root.right)
				is.Equal(bst.Root.left.data, 2)
				is.Equal(bst.Root.right.data, 5)
			},
		},
		{
			name: "insert element to left sub-tree",
			args: args{
				data: 1,
			},
			bst: bst,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				is.EqualValues(bst.nodeCount, 4)

				is.NotNil(bst.Root.left)
				is.NotNil(bst.Root.right)
				is.Nil(bst.Root.left.right)
				is.NotNil(bst.Root.left.left)

				is.Equal(bst.Root.left.left.data, 1)
			},
		},
		{
			name: "insert element to right sub-tree",
			args: args{
				data: 6,
			},
			bst: bst,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				is.EqualValues(bst.nodeCount, 5)

				is.NotNil(bst.Root.left)
				is.NotNil(bst.Root.right)
				is.Nil(bst.Root.left.right)
				is.NotNil(bst.Root.left.left)
				is.NotNil(bst.Root.right.right)

				is.Equal(bst.Root.right.right.data, 6)
			},
		},
		{
			name: "insert one more last element",
			args: args{
				data: 3,
			},
			bst: bst,
			want: func(bst *BinarySearchTree[int], err error) {
				is.Nil(err)
				is.EqualValues(bst.nodeCount, 6)

				is.NotNil(bst.Root.left.right)

				is.Equal(bst.Root.left.right.data, 3)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.bst.Add(tt.args.data)
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
				bst.Add("B")
				bst.Add("A")
				bst.Add("C")
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
