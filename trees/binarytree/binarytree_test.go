package binarytree

import (
	"fmt"
	"testing"

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
				is.Error(err, fmt.Errorf("index 5 is larger than size 4"))
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
				is.Error(err, fmt.Errorf("index 7 is larger than size %d", btree.GetSize()))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			n, err := btree.getNode(tt.index)
			if err != nil {
				tt.want(err)
				return
			}
			err = btree.Delete(n)
			tt.want(err)
		})
	}
}

// func Test_InsertAfter(t *testing.T) {
// 	is := assert.New(t)
// 	bt := NewBinaryTree[int]()

// 	type args struct {
// 		curr *node.Node[int]
// 		data int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		setup func(*node.Node[int])
// 		want  func(error)
// 	}{
// 		{
// 			name: "insert in empty tree",
// 			args: args{
// 				curr: node.NewNode(9),
// 				data: 11,
// 			},
// 			want: func(err error) {
// 				is.Error(fmt.Errorf("empty tree"))
// 				is.Nil(bt.Root)
// 				is.EqualValues(bt.GetSize(), 0)
// 			},
// 		},
// 		{
// 			name: "insert after root node",
// 			args: args{
// 				data: 15,
// 				curr: node.NewNode(10), // root node
// 			},
// 			setup: func(n *node.Node[int]) {
// 				err := bt.insertNode(0, n)
// 				is.Nil(err)
// 			},
// 			want: func(err error) {
// 				is.Nil(err)
// 				is.EqualValues(bt.Root.Data, 10)
// 				is.EqualValues(bt.Root.Right.Data, 15)
// 				is.EqualValues(bt.GetSize(), 2)
// 			},
// 		},
// 		{
// 			name: "insert after node with right child",
// 			args: args{
// 				data: 20,
// 				curr: node.NewNode(11),
// 			},
// 			setup: func(n *node.Node[int]) {
// 				err := bt.insertNode(2, n)
// 				is.Nil(err)
// 			},
// 			want: func(err error) {
// 				is.Nil(err)
// 				is.EqualValues(bt.Root.Data, 10)
// 				is.EqualValues(bt.Root.Right.Data, 15)
// 				is.EqualValues(bt.Root.Right.Left.Data, 20)
// 				is.EqualValues(bt.GetSize(), 3)
// 			},
// 		},
// 		// {
// 		// 	name: "insert after node with successor",
// 		// 	args: args{
// 		// 		data: 12,
// 		// 		curr: 8,
// 		// 	},
// 		// 	setup: func(val int) {
// 		// 		n, err := bt.Insert(val)
// 		// 		is.Nil(err)
// 		// 		is.Equal(n.Data, val)
// 		// 	},
// 		// 	want: func(err error) {
// 		// 		order, err := bt.TraversalOrder(bt.Root)
// 		// 		fmt.Println(order)
// 		// 		is.Nil(err)
// 		// 	},
// 		// },
// 		// {
// 		// 	name: "insert after non-existent node",
// 		// 	args: args{
// 		// 		setup: func(bt *BinaryTree[int], n *node.Node[int]) {
// 		// 			_ = bt.Insert(10)
// 		// 		},
// 		// 		data:    15,
// 		// 		nodeKey: 5,
// 		// 	},
// 		// 	want: func(bt *BinaryTree[int], err error) {
// 		// 		is.Error(err)
// 		// 		is.Equal(err.Error(), "node not found in tree")
// 		// 	},
// 		// },
// 		// Add more test cases as needed
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.setup != nil {
// 				tt.setup(tt.args.curr)
// 			}
// 			err := bt.InsertAfter(tt.args.curr, tt.args.data)
// 			tt.want(err)
// 		})
// 	}
// }

// func Test_TraversalOrder(t *testing.T) {
// 	// is := assert.New(t)

// 	type args struct {
// 		n *node.Node[string]
// 	}

// 	tests := []struct {
// 		name  string
// 		args  args
// 		setup func(*BinaryTree[string])
// 		want  func([]string, error)
// 	}{
// 		{
// 			name: "Test Case 1",
// 			args: args{
// 				n: node.NewNode[string]("alita"),
// 			},
// 			setup: func(bt *BinaryTree[string]) {
// 				// bt.Insert("a")
// 				// bt.Insert("b")
// 				// bt.Insert("c")
// 				// bt.Insert("d")
// 				// bt.Insert("e")
// 				// bt.Insert("f")
// 				// bt.Insert("g")
// 				// bt.Insert("h")
// 				// bt.Insert("i")
// 				// bt.Insert("j")
// 				// bt.Insert("k")
// 				// bt.Insert("l")
// 				// bt.Insert("m")
// 			},
// 			want: func(order []string, err error) {

// 			},
// 		},
// 	}

// 	binaryTree := NewBinaryTree[string]("a", "b", "c", "d", "e")
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.setup != nil {
// 				tt.setup(binaryTree)
// 			}

// 			result, err := binaryTree.TraversalOrder(binaryTree.Root)
// 			tt.want(result, err)
// 		})
// 	}
// }

// func Test_PreOrderTraversal(t *testing.T) {
// 	// is := assert.New(t)

// 	type args struct {
// 		n *node.Node[string]
// 	}

// 	tests := []struct {
// 		name  string
// 		args  args
// 		setup func(*BinaryTree[string])
// 		want  func([]string, error)
// 	}{
// 		{
// 			name: "Test Case 1",
// 			args: args{
// 				n: node.NewNode[string]("alita"),
// 			},
// 			setup: func(bt *BinaryTree[string]) {
// 				// bt.Insert("a")
// 				// bt.Insert("b")
// 				// bt.Insert("c")
// 				// bt.Insert("d")
// 				// bt.Insert("e")
// 				// bt.Insert("f")
// 				// bt.Insert("g")
// 				// bt.Insert("h")
// 				// bt.Insert("i")
// 				// bt.Insert("j")
// 				// bt.Insert("k")
// 				// bt.Insert("l")
// 				// bt.Insert("m")
// 			},
// 			want: func(order []string, err error) {

// 			},
// 		},
// 	}

// 	binaryTree := NewBinaryTree[string]("a", "b", "c", "d", "e")
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.setup != nil {
// 				tt.setup(binaryTree)
// 			}

// 			result, err := binaryTree.PreOrderTraversal(binaryTree.Root)
// 			tt.want(result, err)
// 		})
// 	}
// }

// func Test_PostOrderTraversal(t *testing.T) {
// 	// is := assert.New(t)

// 	type args struct {
// 		n *node.Node[string]
// 	}

// 	tests := []struct {
// 		name  string
// 		args  args
// 		setup func(*BinaryTree[string])
// 		want  func([]string, error)
// 	}{
// 		{
// 			name: "Test Case 1",
// 			args: args{
// 				n: node.NewNode[string]("alita"),
// 			},
// 			setup: func(bt *BinaryTree[string]) {
// 				// bt.Insert("a")
// 				// bt.Insert("b")
// 				// bt.Insert("c")
// 				// bt.Insert("d")
// 				// bt.Insert("e")
// 				// bt.Insert("f")
// 				// bt.Insert("g")
// 				// bt.Insert("h")
// 				// bt.Insert("i")
// 				// bt.Insert("j")
// 				// bt.Insert("k")
// 				// bt.Insert("l")
// 				// bt.Insert("m")
// 			},
// 			want: func(order []string, err error) {

// 			},
// 		},
// 	}

// 	binaryTree := NewBinaryTree[string]("a", "b", "c", "d", "e")
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.setup != nil {
// 				tt.setup(binaryTree)
// 			}

// 			result, err := binaryTree.PostOrderTraversal(binaryTree.Root)
// 			tt.want(result, err)
// 		})
// 	}
// }
