package avltree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	is := assert.New(t)
	avl := new(AVLTree[string, string])

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
				is.EqualValues(1, avl.GetSize())
				is.Equal(avl.Root.GetKey(), "key1")

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(2, avl.GetSize())
				is.Equal(avl.Root.Right.GetKey(), "key2")

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(3, avl.GetSize())
				is.Equal(avl.Root.Left.GetKey(), "key0")

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(3, avl.GetSize())
				is.Equal(avl.Root.Right.GetKey(), "key2")

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(3, avl.GetSize())

				o, err := avl.TraversalOrder(avl.Root)
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
				is.EqualValues(3, avl.GetSize())

				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []string{"value0", "value1", "value2"})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := avl.Insert(tt.key, tt.val)
			tt.want(err)
		})
	}
}

func Test_Delete(t *testing.T) {
	is := assert.New(t)
	avl := new(AVLTree[int, string])
	avl.Insert(10, "10")
	avl.Insert(5, "5")
	avl.Insert(15, "15")

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
				is.EqualValues(avl.GetSize(), 2)
				is.EqualValues(avl.GetHeight(), 1)
				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []string{"10", "15"})
			},
		},
		{
			name: "Delete node with one child",
			key:  10,
			want: func(got string, err error) {
				is.Nil(err)
				is.EqualValues(avl.GetSize(), 1)
				is.EqualValues(avl.GetHeight(), 0)
				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []string{"15"})
			},
		},
		{
			name: "Delete node with two children (predecessor case)",
			setup: func() {
				avl = new(AVLTree[int, string])
				avl.Insert(10, "10")
				avl.Insert(5, "5")
				avl.Insert(15, "15")
			},
			key: 10,
			want: func(got string, err error) {
				is.Nil(err)
				is.EqualValues(avl.GetSize(), 2)
				is.EqualValues(avl.GetHeight(), 1)
				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []string{"5", "15"})
			},
		},
		{
			name: "Delete node with one child (successor case)",
			setup: func() {
				avl = new(AVLTree[int, string])
				avl.Insert(10, "10")
				avl.Insert(15, "15")
			},
			key: 10,
			want: func(got string, err error) {
				is.Nil(err)
				is.EqualValues(avl.GetSize(), 1)
				is.EqualValues(avl.GetHeight(), 0)
				o, err := avl.TraversalOrder(avl.Root)
				is.Nil(err)
				is.Equal(o, []string{"15"})
			},
		},
		{
			name: "Delete non-existent node",
			setup: func() {
				avl = new(AVLTree[int, string])
				avl.Insert(10, "10")
				avl.Insert(5, "5")
				avl.Insert(15, "15")
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
			value, err := avl.Delete(tt.key)
			tt.want(value, err)
		})
	}
}

func Test_TraversalOrder(t *testing.T) {
	is := assert.New(t)
	avl := new(AVLTree[int, string])

	tests := []struct {
		name  string
		setup func()
		want  func([]string, error)
	}{
		{
			name: "simple traversal order",
			setup: func() {
				avl.Insert(1, "a")
				avl.Insert(2, "b")
				avl.Insert(3, "c")
				avl.Insert(4, "d")
				avl.Insert(5, "e")
				avl.Insert(6, "f")
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
			order, err := avl.TraversalOrder(avl.Root)
			tt.want(order, err)
		})
	}
}

func Test_PreOrderTraversal(t *testing.T) {
	is := assert.New(t)
	avl := new(AVLTree[int, string])

	tests := []struct {
		name  string
		setup func()
		want  func([]string, error)
	}{
		{
			name: "simple traversal order",
			setup: func() {
				avl.Insert(1, "a")
				avl.Insert(2, "b")
				avl.Insert(3, "c")
				avl.Insert(4, "d")
				avl.Insert(5, "e")
				avl.Insert(6, "f")
			},
			want: func(order []string, err error) {
				is.Nil(err)
				is.Equal(order, []string{"d", "b", "a", "c", "e", "f"})
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
	avl := new(AVLTree[int, string])

	tests := []struct {
		name  string
		setup func()
		want  func([]string, error)
	}{
		{
			name: "simple traversal order",
			setup: func() {
				avl.Insert(1, "a")
				avl.Insert(2, "b")
				avl.Insert(3, "c")
				avl.Insert(4, "d")
				avl.Insert(5, "e")
				avl.Insert(6, "f")
			},
			want: func(order []string, err error) {
				is.Nil(err)
				is.Equal(order, []string{"a", "c", "b", "f", "e", "d"})
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
