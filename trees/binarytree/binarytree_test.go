package binarytree

import (
	"testing"

	"github.com/OladapoAjala/datastructures/trees/node"
)

func Test_TraversalOrder(t *testing.T) {
	// is := assert.New(t)

	type args struct {
		n *node.Node[string]
	}

	tests := []struct {
		name  string
		args  args
		setup func(*BinaryTree[string])
		want  func([]string, error)
	}{
		{
			name: "Test Case 1",
			args: args{
				n: node.NewNode[string]("alita"),
			},
			setup: func(bt *BinaryTree[string]) {
				bt.Insert("a")
				bt.Insert("b")
				bt.Insert("c")
				bt.Insert("d")
				bt.Insert("e")
				// bt.Insert("f")
				// bt.Insert("g")
				// bt.Insert("h")
				// bt.Insert("i")
				// bt.Insert("j")
				// bt.Insert("k")
				// bt.Insert("l")
				// bt.Insert("m")
			},
			want: func(order []string, err error) {

			},
		},
	}

	binaryTree := NewBinaryTree[string]()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(binaryTree)
			}

			result, err := binaryTree.TraversalOrder(binaryTree.Root)
			tt.want(result, err)
		})
	}
}

func Test_PreOrderTraversal(t *testing.T) {
	// is := assert.New(t)

	type args struct {
		n *node.Node[string]
	}

	tests := []struct {
		name  string
		args  args
		setup func(*BinaryTree[string])
		want  func([]string, error)
	}{
		{
			name: "Test Case 1",
			args: args{
				n: node.NewNode[string]("alita"),
			},
			setup: func(bt *BinaryTree[string]) {
				bt.Insert("a")
				bt.Insert("b")
				bt.Insert("c")
				bt.Insert("d")
				bt.Insert("e")
				// bt.Insert("f")
				// bt.Insert("g")
				// bt.Insert("h")
				// bt.Insert("i")
				// bt.Insert("j")
				// bt.Insert("k")
				// bt.Insert("l")
				// bt.Insert("m")
			},
			want: func(order []string, err error) {

			},
		},
	}

	binaryTree := NewBinaryTree[string]()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(binaryTree)
			}

			result, err := binaryTree.PreOrderTraversal(binaryTree.Root)
			tt.want(result, err)
		})
	}
}

func Test_PostOrderTraversal(t *testing.T) {
	// is := assert.New(t)

	type args struct {
		n *node.Node[string]
	}

	tests := []struct {
		name  string
		args  args
		setup func(*BinaryTree[string])
		want  func([]string, error)
	}{
		{
			name: "Test Case 1",
			args: args{
				n: node.NewNode[string]("alita"),
			},
			setup: func(bt *BinaryTree[string]) {
				bt.Insert("a")
				bt.Insert("b")
				bt.Insert("c")
				// bt.Insert("d")
				// bt.Insert("e")
				// bt.Insert("f")
				// bt.Insert("g")
				// bt.Insert("h")
				// bt.Insert("i")
				// bt.Insert("j")
				// bt.Insert("k")
				// bt.Insert("l")
				// bt.Insert("m")
			},
			want: func(order []string, err error) {

			},
		},
	}

	binaryTree := NewBinaryTree[string]()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(binaryTree)
			}

			result, err := binaryTree.PostOrderTraversal(binaryTree.Root)
			tt.want(result, err)
		})
	}
}
