package binarysearchtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	is := assert.New(t)
	bst := new(BinarySearchTree[string, string])

	tests := []struct {
		name string
		key  string
		val  string
		want func(error)
	}{
		{
			name: "Insert into an empty tree",
			key:  "key1",
			val:  "value1",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(1, bst.GetSize())
				is.Equal(bst.Root.GetKey(), "key1")

				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"value1"})
			},
		},
		{
			name: "Insert at the end",
			key:  "key2",
			val:  "value2",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(2, bst.GetSize())
				is.Equal(bst.Root.Right.GetKey(), "key2")

				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"value1", "value2"})
			},
		},
		{
			name: "Insert at the beginning",
			key:  "key0",
			val:  "value0",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(3, bst.GetSize())
				is.Equal(bst.Root.Left.GetKey(), "key0")

				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"value0", "value1", "value2"})
			},
		},
		{
			name: "Insert at the end with same key",
			key:  "key2",
			val:  "new_value",
			want: func(err error) {
				is.Error(err, fmt.Errorf("key key2 already in tree"))
				is.EqualValues(3, bst.GetSize())
				is.Equal(bst.Root.Right.GetKey(), "key2")

				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"value0", "value1", "value2"})
			},
		},
		{
			name: "Insert with empty key",
			key:  "",
			val:  "empty_value",
			want: func(err error) {
				is.Error(err, fmt.Errorf("empty key"))
				is.EqualValues(3, bst.GetSize())

				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"value0", "value1", "value2"})
			},
		},
		{
			name: "Insert with empty value",
			key:  "key3",
			val:  "",
			want: func(err error) {
				is.Error(err, fmt.Errorf("empty value"))
				is.EqualValues(3, bst.GetSize())

				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"value0", "value1", "value2"})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := bst.Insert(tt.key, tt.val)
			tt.want(err)
		})
	}
}

func Test_Delete(t *testing.T) {
	is := assert.New(t)
	bst := new(BinarySearchTree[int, string])
	bst.Insert(10, "10")
	bst.Insert(5, "5")
	bst.Insert(15, "15")

	tests := []struct {
		name  string
		setup func()
		key   int
		want  func(string, error)
	}{
		{
			name: "Delete leaf node",
			key:  5,
			want: func(got string, err error) {
				is.Nil(err)
				is.EqualValues(bst.GetSize(), 2)
				is.EqualValues(bst.GetHeight(), 1)
				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"10", "15"})
			},
		},
		{
			name: "Delete node with one child",
			key:  10,
			want: func(got string, err error) {
				is.Nil(err)
				is.EqualValues(bst.GetSize(), 1)
				is.EqualValues(bst.GetHeight(), 0)
				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"15"})
			},
		},
		{
			name: "Delete node with two children (predecessor case)",
			setup: func() {
				bst = new(BinarySearchTree[int, string])
				bst.Insert(10, "10")
				bst.Insert(5, "5")
				bst.Insert(15, "15")
			},
			key: 10,
			want: func(got string, err error) {
				is.Nil(err)
				is.EqualValues(bst.GetSize(), 2)
				is.EqualValues(bst.GetHeight(), 1)
				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"5", "15"})
			},
		},
		{
			name: "Delete node with one child (successor case)",
			setup: func() {
				bst = new(BinarySearchTree[int, string])
				bst.Insert(10, "10")
				bst.Insert(15, "15")
			},
			key: 10,
			want: func(got string, err error) {
				is.Nil(err)
				is.EqualValues(bst.GetSize(), 1)
				is.EqualValues(bst.GetHeight(), 0)
				o, err := bst.TraversalOrder(bst.Root)
				is.Nil(err)
				is.Equal(o, []string{"15"})
			},
		},
		{
			name: "Delete non-existent node",
			setup: func() {
				bst = new(BinarySearchTree[int, string])
				bst.Insert(10, "10")
				bst.Insert(5, "5")
				bst.Insert(15, "15")
			},
			key: 7,
			want: func(got string, err error) {
				is.Error(err, fmt.Errorf("key 7 is not in tree"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			value, err := bst.Delete(tt.key)
			tt.want(value, err)
		})
	}
}

func Test_Find(t *testing.T) {
	is := assert.New(t)

	type args struct {
		key int
	}
	tests := []struct {
		name  string
		args  args
		setup func() *BinarySearchTree[int, string]
		want  func(string, error)
	}{
		{
			name: "find random (key 5) element",
			args: args{
				key: 5,
			},
			setup: func() *BinarySearchTree[int, string] {
				bst := new(BinarySearchTree[int, string])
				bst.Insert(10, "10")
				bst.Insert(5, "5")
				bst.Insert(15, "15")
				return bst
			},
			want: func(got string, err error) {
				is.Nil(err)
				is.Equal(got, "5")
			},
		},
		{
			name: "find in empty array",
			args: args{
				key: 0,
			},
			setup: func() *BinarySearchTree[int, string] {
				return new(BinarySearchTree[int, string])
			},
			want: func(got string, err error) {
				is.NotNil(err)
				is.Error(err, fmt.Errorf("empty tree"))
				is.Empty(got)
			},
		},
		{
			name: "key less than min",
			args: args{
				key: -1,
			},
			setup: func() *BinarySearchTree[int, string] {
				bst := new(BinarySearchTree[int, string])
				bst.Insert(10, "10")
				return bst
			},
			want: func(got string, err error) {
				is.NotNil(err)
				is.Error(err, fmt.Errorf("key -1 is not in tree"))
			},
		},
		{
			name: "key greater than max",
			args: args{
				key: 12,
			},
			setup: func() *BinarySearchTree[int, string] {
				bst := new(BinarySearchTree[int, string])
				bst.Insert(10, "10")
				return bst
			},
			want: func(got string, err error) {
				is.NotNil(err)
				is.Error(err, fmt.Errorf("key 2 is not in tree"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := tt.setup()
			data, err := bst.Find(tt.args.key)
			tt.want(data, err)
		})
	}
}

func Test_TraversalOrder(t *testing.T) {
	is := assert.New(t)
	bst := new(BinarySearchTree[int, string])

	tests := []struct {
		name  string
		setup func()
		want  func([]string, error)
	}{
		{
			name: "simple traversal order",
			setup: func() {
				bst.Insert(3, "c")
				bst.Insert(1, "a")
				bst.Insert(2, "b")
				bst.Insert(4, "d")
				bst.Insert(5, "e")
				bst.Insert(6, "f")
			},
			want: func(order []string, err error) {
				is.Nil(err)
				is.Equal(order, []string{"a", "b", "c", "d", "e", "f"})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			order, err := bst.TraversalOrder(bst.Root)
			tt.want(order, err)
		})
	}
}

func Test_PreOrderTraversal(t *testing.T) {
	is := assert.New(t)
	bst := new(BinarySearchTree[int, string])

	tests := []struct {
		name  string
		setup func()
		want  func([]string, error)
	}{
		{
			name: "simple pre-order order",
			setup: func() {
				bst.Insert(3, "c")
				bst.Insert(1, "a")
				bst.Insert(2, "b")
				bst.Insert(4, "d")
				bst.Insert(5, "e")
				bst.Insert(6, "f")
			},
			want: func(order []string, err error) {
				is.Nil(err)
				is.Equal(order, []string{"c", "a", "b", "d", "e", "f"})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			order, err := bst.PreOrderTraversal(bst.Root)
			tt.want(order, err)
		})
	}
}

func Test_PostOrderTraversal(t *testing.T) {
	is := assert.New(t)
	bst := new(BinarySearchTree[int, string])

	tests := []struct {
		name  string
		setup func()
		want  func([]string, error)
	}{
		{
			name: "simple post-order order",
			setup: func() {
				bst.Insert(3, "c")
				bst.Insert(1, "a")
				bst.Insert(2, "b")
				bst.Insert(4, "d")
				bst.Insert(5, "e")
				bst.Insert(6, "f")
			},
			want: func(order []string, err error) {
				is.Nil(err)
				is.Equal(order, []string{"b", "a", "f", "e", "d", "c"})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			order, err := bst.PostOrderTraversal(bst.Root)
			tt.want(order, err)
		})
	}
}
