package binarytree

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/trees/node"
	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[int]()

	tests := []struct {
		name  string
		index int32
		data  int
		want  func(error)
	}{
		{
			name:  "Insert into an empty tree",
			index: 0,
			data:  10,
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(1, btree.GetSize())
				is.Equal(btree.Root.Data, 10)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{10})
			},
		},
		{
			name:  "Insert at the end",
			index: 1,
			data:  20,
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(2, btree.GetSize())
				is.Equal(btree.Root.Right.Data, 20)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 20})
			},
		},
		{
			name:  "Insert at a specific index",
			index: 1,
			data:  15,
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(3, btree.GetSize())
				is.Equal(btree.Root.Data, 10)
				is.Equal(btree.Root.Right.Data, 20)
				is.Equal(btree.Root.Right.Left.Data, 15)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 15, 20})
			},
		},
		{
			name:  "Insert at the beginning",
			index: 0,
			data:  5,
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(4, btree.GetSize())
				is.Equal(btree.Root.Left.Data, 5)
				is.Equal(btree.Root.Data, 10)
				is.Equal(btree.Root.Right.Data, 20)
				is.Equal(btree.Root.Right.Left.Data, 15)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{5, 10, 15, 20})
			},
		},
		{
			name:  "Insert at an invalid index",
			index: 5,
			data:  25,
			want: func(err error) {
				is.Error(err, fmt.Errorf("index 5 is out of range"))
				is.EqualValues(4, btree.GetSize())
				is.Equal(btree.Root.Left.Data, 5)
				is.Equal(btree.Root.Data, 10)
				is.Equal(btree.Root.Right.Data, 20)
				is.Equal(btree.Root.Right.Left.Data, 15)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{5, 10, 15, 20})
			},
		},
		{
			name:  "Insert with empty data",
			index: 2,
			data:  0,
			want: func(err error) {
				is.Error(err, fmt.Errorf("empty data"))
				is.EqualValues(4, btree.GetSize())
				is.Equal(btree.Root.Left.Data, 5)
				is.Equal(btree.Root.Data, 10)
				is.Equal(btree.Root.Right.Data, 20)
				is.Equal(btree.Root.Right.Left.Data, 15)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{5, 10, 15, 20})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := btree.Insert(tt.index, tt.data)
			tt.want(err)
		})
	}
}

func Test_Delete(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[int](10, 5, 15)

	tests := []struct {
		name  string
		setup func()
		index int32
		want  func(error)
	}{
		{
			name:  "Delete leaf node",
			index: 2,
			want: func(err error) {
				is.Nil(err)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 5})
			},
		},
		{
			name:  "Delete node with one child",
			index: 0,
			want: func(err error) {
				is.Nil(err)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{5})
			},
		},
		{
			name: "Delete node with two children (predecessor case)",
			setup: func() {
				btree = NewBinaryTree[int](10, 5, 15, 3, 8, 12, 20)
			},
			index: 2,
			want: func(err error) {
				is.Nil(err)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 5, 3, 8, 12, 20})
			},
		},
		{
			name: "Delete node with two children (successor case)",
			setup: func() {
				btree = NewBinaryTree[int](10, 5, 15, 3, 8, 12, 20)
			},
			index: 1,
			want: func(err error) {
				is.Nil(err)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 15, 3, 8, 12, 20})
			},
		},
		{
			name: "Delete root node with two children",
			setup: func() {
				btree = NewBinaryTree[int](10, 5, 15, 3, 8, 12, 20)
			},
			index: 0,
			want: func(err error) {
				is.Nil(err)

				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{5, 15, 3, 8, 12, 20})
			},
		},
		{
			name:  "Delete non-existent node",
			index: 7,
			want: func(err error) {
				is.Error(err, fmt.Errorf("index 7 is out of range"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			err := btree.Delete(tt.index)
			tt.want(err)
		})
	}
}

func Test_InsertAfter(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[int](10)

	type args struct {
		currIndex int32
		data      *node.Node[int]
	}
	tests := []struct {
		name  string
		args  args
		setup func()
		want  func(error)
	}{
		{
			name: "insert after root node",
			args: args{
				data:      node.NewNode[int](15),
				currIndex: 0,
			},
			want: func(err error) {
				is.Nil(err)
				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 15})
			},
		},
		{
			name: "insert after node with right child",
			args: args{
				data:      node.NewNode[int](20),
				currIndex: 0,
			},
			want: func(err error) {
				is.Nil(err)
				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 20, 15})
			},
		},
		{
			name: "insert after node with an ancestor successor",
			args: args{
				data:      node.NewNode[int](13),
				currIndex: 1,
			},
			setup: func() {
				n, err := btree.getNode(1)
				is.Nil(err)
				err = btree.InsertBefore(n, node.NewNode[int](17))
				is.Nil(err)
			},
			want: func(err error) {
				is.Nil(err)
				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 17, 13, 20, 15})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			n, err := btree.getNode(tt.args.currIndex)
			is.Nil(err)
			err = btree.InsertAfter(n, tt.args.data)
			tt.want(err)
		})
	}
}

func Test_InsertBefore(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[int](10)

	type args struct {
		currIndex int32
		data      *node.Node[int]
	}
	tests := []struct {
		name  string
		args  args
		setup func()
		want  func(error)
	}{
		{
			name: "insert before root node",
			args: args{
				data:      node.NewNode[int](5),
				currIndex: 0,
			},
			want: func(err error) {
				is.Nil(err)
				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{5, 10})
			},
		},
		{
			name: "insert before node with left child",
			args: args{
				data:      node.NewNode[int](7),
				currIndex: 1,
			},
			want: func(err error) {
				is.Nil(err)
				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{5, 7, 10})
			},
		},
		{
			name: "insert before node with ancestor predecessor",
			args: args{
				data:      node.NewNode[int](12),
				currIndex: 1,
			},
			setup: func() {
				n, err := btree.getNode(0)
				is.Nil(err)
				err = btree.InsertAfter(n, node.NewNode[int](8))
				is.Nil(err)
			},
			want: func(err error) {
				is.Nil(err)
				o, err := btree.TraversalOrder(btree.Root)
				is.Nil(err)
				is.Equal(o, []int{5, 12, 8, 7, 10})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			n, err := btree.getNode(tt.args.currIndex)
			is.Nil(err)
			err = btree.InsertBefore(n, tt.args.data)
			tt.want(err)
		})
	}
}

func Test_Contains(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[string]()
	btree.Insert(0, "a")
	btree.Insert(1, "b")
	btree.Insert(0, "c")
	btree.Insert(1, "d")
	btree.Insert(0, "e")
	btree.Insert(4, "f")
	btree.Insert(3, "g")
	is.True(btree.Contains("e"))
	is.False(btree.Contains("j"))
}

func Test_InsertFirst(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[string]()
	btree.Insert(0, "a")
	btree.Insert(1, "b")
	btree.Insert(0, "c")
	btree.Insert(1, "d")
	btree.Insert(0, "e")
	btree.Insert(4, "f")
	btree.Insert(3, "g")

	btree.InsertFirst("z")
	o, err := btree.TraversalOrder(btree.Root)
	is.Nil(err)
	is.Equal(o, []string{"z", "e", "c", "d", "g", "a", "f", "b"})
}

func Test_InsertLast(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[string]()
	btree.Insert(0, "a")
	btree.Insert(1, "b")
	btree.Insert(0, "c")
	btree.Insert(1, "d")
	btree.Insert(0, "e")
	btree.Insert(4, "f")
	btree.Insert(3, "g")

	btree.InsertLast("z")
	o, err := btree.TraversalOrder(btree.Root)
	is.Nil(err)
	is.Equal(o, []string{"e", "c", "d", "g", "a", "f", "b", "z"})
}

func Test_DeleteFirst(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name string
		bt   *BinaryTree[string]
		want func(*BinaryTree[string], error)
	}{
		{
			name: "simple delete first",
			bt:   NewBinaryTree("Node 0", "Node 1"),
			want: func(bt *BinaryTree[string], err error) {
				is.Nil(err)
				is.False(bt.Contains("Node 0"))

				data, err := bt.GetData(0)
				is.Nil(err)
				is.Equal(data, "Node 1")
			},
		},
		{
			name: "delete first (only node)",
			bt:   NewBinaryTree("A"),
			want: func(bt *BinaryTree[string], err error) {
				is.Nil(err)
				is.False(bt.Contains("A"))

				data, err := bt.GetData(0)
				is.Empty(data)
				is.Error(err, fmt.Errorf("index 0 is out of range"))
			},
		},
		{
			name: "delete first (empty node)",
			bt:   NewBinaryTree[string](),
			want: func(bt *BinaryTree[string], err error) {
				is.Error(err, fmt.Errorf("cannot delete from empty tree"))

				data, err := bt.GetData(0)
				is.Empty(data)
				is.Error(err, fmt.Errorf("index 0 is out of range"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.bt.DeleteFirst()
			tt.want(tt.bt, err)
		})
	}
}

func Test_TraversalOrder(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[string]()

	type args struct {
		n *node.Node[string]
	}

	tests := []struct {
		name  string
		args  args
		setup func()
		want  func([]string, error)
	}{
		{
			name: "Test Case 1",
			args: args{
				n: node.NewNode[string]("alita"),
			},
			setup: func() {
				btree.Insert(0, "a")
				btree.Insert(1, "b")
				btree.Insert(0, "c")
				btree.Insert(1, "d")
				btree.Insert(0, "e")
				btree.Insert(4, "f")
				btree.Insert(3, "g")
			},
			want: func(order []string, err error) {
				is.Nil(err)
				is.Equal(order, []string{"e", "c", "d", "g", "a", "f", "b"})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			order, err := btree.TraversalOrder(btree.Root)
			tt.want(order, err)
		})
	}
}

func Test_PreOrderTraversal(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[string]()

	type args struct {
		n *node.Node[string]
	}

	tests := []struct {
		name  string
		args  args
		setup func()
		want  func([]string, error)
	}{
		{
			name: "Test Case 1",
			args: args{
				n: node.NewNode[string]("alita"),
			},
			setup: func() {
				btree.Insert(0, "a")
				btree.Insert(1, "b")
				btree.Insert(0, "c")
				btree.Insert(1, "d")
				btree.Insert(0, "e")
				btree.Insert(4, "f")
				btree.Insert(3, "g")
			},
			want: func(order []string, err error) {
				is.Nil(err)
				is.Equal(order, []string{"a", "c", "e", "d", "g", "b", "f"})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			order, err := btree.PreOrderTraversal(btree.Root)
			tt.want(order, err)
		})
	}
}

func Test_PostOrderTraversal(t *testing.T) {
	is := assert.New(t)
	btree := NewBinaryTree[string]()

	type args struct {
		n *node.Node[string]
	}

	tests := []struct {
		name  string
		args  args
		setup func()
		want  func([]string, error)
	}{
		{
			name: "Test Case 1",
			args: args{
				n: node.NewNode[string]("alita"),
			},
			setup: func() {
				btree.Insert(0, "a")
				btree.Insert(1, "b")
				btree.Insert(0, "c")
				btree.Insert(1, "d")
				btree.Insert(0, "e")
				btree.Insert(4, "f")
				btree.Insert(3, "g")
			},
			want: func(order []string, err error) {
				is.Nil(err)
				is.Equal(order, []string{"e", "g", "d", "c", "f", "b", "a"})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			order, err := btree.PostOrderTraversal(btree.Root)
			tt.want(order, err)
		})
	}
}
