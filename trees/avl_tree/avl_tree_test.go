package avltree

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/trees/node"
	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[int]()

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
				is.EqualValues(1, avl.GetSize())
				is.Equal(avl.Root.Data, 10)

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(2, avl.GetSize())
				is.Equal(avl.Root.Right.Data, 20)

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(3, avl.GetSize())
				is.Equal(avl.Root.Data, 10)
				is.Equal(avl.Root.Right.Data, 20)
				is.Equal(avl.Root.Right.Left.Data, 15)

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(4, avl.GetSize())
				is.Equal(avl.Root.Left.Data, 5)
				is.Equal(avl.Root.Data, 10)
				is.Equal(avl.Root.Right.Data, 20)
				is.Equal(avl.Root.Right.Left.Data, 15)

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(4, avl.GetSize())
				is.Equal(avl.Root.Left.Data, 5)
				is.Equal(avl.Root.Data, 10)
				is.Equal(avl.Root.Right.Data, 20)
				is.Equal(avl.Root.Right.Left.Data, 15)

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(4, avl.GetSize())
				is.Equal(avl.Root.Left.Data, 5)
				is.Equal(avl.Root.Data, 10)
				is.Equal(avl.Root.Right.Data, 20)
				is.Equal(avl.Root.Right.Left.Data, 15)

				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []int{5, 10, 15, 20})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := avl.Insert(tt.index, tt.data)
			tt.want(err)
		})
	}
}

func Test_Delete(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[int](10, 5, 15)

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

				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 5})
			},
		},
		{
			name:  "Delete node with one child",
			index: 0,
			want: func(err error) {
				is.Nil(err)

				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []int{5})
			},
		},
		{
			name: "Delete node with two children (predecessor case)",
			setup: func() {
				avl = NewAVLTree[int](10, 5, 15, 3, 8, 12, 20)
			},
			index: 2,
			want: func(err error) {
				is.Nil(err)

				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 5, 3, 8, 12, 20})
			},
		},
		{
			name: "Delete node with two children (successor case)",
			setup: func() {
				avl = NewAVLTree[int](10, 5, 15, 3, 8, 12, 20)
			},
			index: 1,
			want: func(err error) {
				is.Nil(err)

				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []int{10, 15, 3, 8, 12, 20})
			},
		},
		{
			name: "Delete root node with two children",
			setup: func() {
				avl = NewAVLTree[int](10, 5, 15, 3, 8, 12, 20)
			},
			index: 0,
			want: func(err error) {
				is.Nil(err)

				o, err := avl.TraversalOrder(avl.Root)
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
			err := avl.Delete(tt.index)
			tt.want(err)
		})
	}
}

func Test_InsertAfter(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[int](10)

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
				o, err := avl.TraversalOrder(avl.Root)
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
				o, err := avl.TraversalOrder(avl.Root)
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
				n, err := avl.getNode(1)
				is.Nil(err)
				err = avl.InsertBefore(n, node.NewNode[int](17))
				is.Nil(err)
			},
			want: func(err error) {
				is.Nil(err)
				o, err := avl.TraversalOrder(avl.Root)
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

			n, err := avl.getNode(tt.args.currIndex)
			is.Nil(err)
			err = avl.InsertAfter(n, tt.args.data)
			tt.want(err)
		})
	}
}

func Test_InsertBefore(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[int](10)

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
				o, err := avl.TraversalOrder(avl.Root)
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
				o, err := avl.TraversalOrder(avl.Root)
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
				n, err := avl.getNode(0)
				is.Nil(err)
				err = avl.InsertAfter(n, node.NewNode[int](8))
				is.Nil(err)
			},
			want: func(err error) {
				is.Nil(err)
				o, err := avl.TraversalOrder(avl.Root)
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

			n, err := avl.getNode(tt.args.currIndex)
			is.Nil(err)
			err = avl.InsertBefore(n, tt.args.data)
			tt.want(err)
		})
	}
}

func Test_Contains(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[string]()
	avl.Insert(0, "a")
	avl.Insert(1, "b")
	avl.Insert(0, "c")
	avl.Insert(1, "d")
	avl.Insert(0, "e")
	avl.Insert(4, "f")
	avl.Insert(3, "g")
	is.True(avl.Contains("e"))
	is.False(avl.Contains("j"))
}

func Test_InsertFirst(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[string]()
	avl.Insert(0, "a")
	avl.Insert(1, "b")
	avl.Insert(0, "c")
	avl.Insert(1, "d")
	avl.Insert(0, "e")
	avl.Insert(4, "f")
	avl.Insert(3, "g")

	avl.InsertFirst("z")
	o, err := avl.TraversalOrder(avl.Root)
	is.Nil(err)
	is.Equal(o, []string{"z", "e", "c", "d", "g", "a", "f", "b"})
}

func Test_InsertLast(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[string]()
	avl.Insert(0, "a")
	avl.Insert(1, "b")
	avl.Insert(0, "c")
	avl.Insert(1, "d")
	avl.Insert(0, "e")
	avl.Insert(4, "f")
	avl.Insert(3, "g")

	avl.InsertLast("z")
	o, err := avl.TraversalOrder(avl.Root)
	is.Nil(err)
	is.Equal(o, []string{"e", "c", "d", "g", "a", "f", "b", "z"})
}

func Test_DeleteFirst(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name string
		avl  *AVLTree[string]
		want func(*AVLTree[string], error)
	}{
		{
			name: "simple delete first",
			avl:  NewAVLTree("Node 0", "Node 1"),
			want: func(avl *AVLTree[string], err error) {
				is.Nil(err)
				is.False(avl.Contains("Node 0"))

				data, err := avl.GetData(0)
				is.Nil(err)
				is.Equal(data, "Node 1")
			},
		},
		{
			name: "delete first (only node)",
			avl:  NewAVLTree("A"),
			want: func(avl *AVLTree[string], err error) {
				is.Nil(err)
				is.False(avl.Contains("A"))

				data, err := avl.GetData(0)
				is.Empty(data)
				is.Error(err, fmt.Errorf("index 0 is out of range"))
			},
		},
		{
			name: "delete first (empty node)",
			avl:  NewAVLTree[string](),
			want: func(avl *AVLTree[string], err error) {
				is.Error(err, fmt.Errorf("cannot delete from empty tree"))

				data, err := avl.GetData(0)
				is.Empty(data)
				is.Error(err, fmt.Errorf("index 0 is out of range"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.avl.DeleteFirst()
			tt.want(tt.avl, err)
		})
	}
}

func Test_TraversalOrder(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[string]()

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
				avl.Insert(0, "a")
				avl.Insert(1, "b")
				avl.Insert(0, "c")
				avl.Insert(1, "d")
				avl.Insert(0, "e")
				avl.Insert(4, "f")
				avl.Insert(3, "g")
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
			order, err := avl.TraversalOrder(avl.Root)
			tt.want(order, err)
		})
	}
}

func Test_PreOrderTraversal(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[string]()

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
				avl.Insert(0, "a")
				avl.Insert(1, "b")
				avl.Insert(0, "c")
				avl.Insert(1, "d")
				avl.Insert(0, "e")
				avl.Insert(4, "f")
				avl.Insert(3, "g")
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
			order, err := avl.PreOrderTraversal(avl.Root)
			tt.want(order, err)
		})
	}
}

func Test_PostOrderTraversal(t *testing.T) {
	is := assert.New(t)
	avl := NewAVLTree[string]()

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
				avl.Insert(0, "a")
				avl.Insert(1, "b")
				avl.Insert(0, "c")
				avl.Insert(1, "d")
				avl.Insert(0, "e")
				avl.Insert(4, "f")
				avl.Insert(3, "g")
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
			order, err := avl.PostOrderTraversal(avl.Root)
			tt.want(order, err)
		})
	}
}
