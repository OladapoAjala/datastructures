package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Get(t *testing.T) {
	is := assert.New(t)
	testList := NewList("Node 0", "Node 1", "Node 2")
	type args struct {
		index int32
	}
	tests := []struct {
		name string
		list *LinkedList[string]
		args args
		want func(string, error)
	}{
		{
			name: "get data from empty list",
			list: new(LinkedList[string]),
			args: args{
				index: 0,
			},
			want: func(data string, err error) {
				is.Error(err, "data not found")
			},
		},
		{
			name: "get Node 1 from list",
			list: testList,
			args: args{
				index: 1,
			},
			want: func(data string, err error) {
				is.Nil(err)
				is.Equal(data, "Node 1")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.list.GetData(tt.args.index)
			tt.want(data, err)
		})
	}
}

func TestLinkedList_InsertLast(t *testing.T) {
	testList := NewList("Node 0", "Node 1", "Node 2")

	type args struct {
		data string
	}
	tests := []struct {
		name string
		list *LinkedList[string]
		args args
		want func(*LinkedList[string]) bool
	}{
		{
			name: "append to an empty list",
			list: new(LinkedList[string]),
			args: args{
				data: "DSA sucks",
			},
			want: func(l *LinkedList[string]) bool {
				return l.Head.Data == "DSA sucks" && l.Tail.Data == "DSA sucks"

			},
		},
		{
			name: "append to a list with three nodes",
			list: testList,
			args: args{
				data: "Node 3",
			},
			want: func(l *LinkedList[string]) bool {
				return l.Tail.Data == "Node 3" && l.Head.Data == "Node 0"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.InsertLast(tt.args.data)
			if !tt.want(tt.list) {
				t.Errorf("append failed")
			}
		})
	}
}

func TestLinkedList_InsertFirst(t *testing.T) {
	testList := NewList("Node 0", "Node 1", "Node 2")

	type args struct {
		data string
	}
	tests := []struct {
		name string
		list *LinkedList[string]
		args args
		want func(*LinkedList[string]) bool
	}{
		{
			name: "prepend to an empty list",
			list: new(LinkedList[string]),
			args: args{
				data: "DSA Sucks",
			},
			want: func(l *LinkedList[string]) bool {
				return l.Head.Data == "DSA Sucks" && l.Tail.Data == "DSA Sucks"
			},
		},
		{
			name: "prepend to a list with three nodes",
			list: testList,
			args: args{
				data: "Node 3",
			},
			want: func(l *LinkedList[string]) bool {
				return l.Head.Data == "Node 3" && l.Tail.Data == "Node 2"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.InsertFirst(tt.args.data)
			if !tt.want(tt.list) {
				t.Errorf("prepend failed")
			}
		})
	}
}

func TestLinkedList_Insert(t *testing.T) {
	testList := NewList("Node 0", "Node 1")

	type args struct {
		index int32
		data  string
	}
	tests := []struct {
		name string
		list *LinkedList[string]
		args args
		want func(*LinkedList[string]) bool
	}{
		{
			name: "insert into a list with two nodes",
			list: testList,
			args: args{
				index: 1,
				data:  "Node 2",
			},
			want: func(l *LinkedList[string]) bool {
				return l.Head.Next.Data == "Node 2"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Insert(tt.args.index, tt.args.data)

			if !tt.want(tt.list) {
				t.Errorf("Error inserting into list")
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	is := assert.New(t)

	type args struct {
		index int32
	}
	tests := []struct {
		name string
		args args
		list *LinkedList[string]
		want func(*LinkedList[string], error)
	}{
		{
			name: "remove head",
			args: args{
				index: 0,
			},
			list: NewList("Node 0", "Node 1"),
			want: func(ll *LinkedList[string], err error) {
				is.Nil(err)
				is.EqualValues(ll.Size(), 1)
				is.False(ll.Contains("Node 0"))

				node, err := ll.GetNode(0)
				is.Nil(err)
				is.Equal(node.Data, "Node 1")
			},
		},
		{
			name: "remove tail",
			args: args{
				index: 2,
			},
			list: NewList("A", "B", "C"),
			want: func(ll *LinkedList[string], err error) {
				is.Nil(err)
				is.EqualValues(ll.Size(), 2)
				is.False(ll.Contains("C"))

				node, err := ll.GetNode(2)
				is.Nil(node)
				is.Error(err, "node not found")
			},
		},
		{
			name: "remove from list with three nodes",
			args: args{
				index: 1,
			},
			list: NewList("Node 0", "Node 1", "Node 2"),
			want: func(ll *LinkedList[string], err error) {
				is.Nil(err)
				is.EqualValues(ll.Size(), 2)
				is.False(ll.Contains("Node 1"))

				node, err := ll.GetNode(1)
				is.Nil(err)
				is.Equal(node.Data, "Node 2")
			},
		},
		{
			name: "remove only element in the list",
			args: args{
				index: 0,
			},
			list: NewList("Node 0"),
			want: func(ll *LinkedList[string], err error) {
				is.Nil(err)
				is.Nil(ll.Head)
				is.Nil(ll.Tail)
				is.EqualValues(ll.Length, 0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.list.Delete(tt.args.index)
			tt.want(tt.list, err)
		})
	}
}

func TestLinkedList_DeleteFirst(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name string
		list *LinkedList[string]
		want func(*LinkedList[string], error)
	}{
		{
			name: "simple delete first",
			list: NewList("Node 0", "Node 1"),
			want: func(ll *LinkedList[string], err error) {
				is.Nil(err)
				is.EqualValues(ll.Size(), 1)
				is.False(ll.Contains("Node 0"))

				data, err := ll.GetData(0)
				is.Nil(err)
				is.Equal(data, "Node 1")
			},
		},
		{
			name: "delete first (only node)",
			list: NewList("A"),
			want: func(ll *LinkedList[string], err error) {
				is.Nil(err)
				is.EqualValues(ll.Size(), 0)
				is.False(ll.Contains("A"))

				data, err := ll.GetData(0)
				is.Empty(data)
				is.Error(err, "data not found")
			},
		},
		{
			name: "delete first (empty node)",
			list: NewList[string](),
			want: func(ll *LinkedList[string], err error) {
				is.Error(err, "cannot remove from empty list")
				is.EqualValues(ll.Size(), 0)

				node, err := ll.GetNode(0)
				is.Nil(node)
				is.Error(err, "node not found")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.list.DeleteFirst()
			tt.want(tt.list, err)
		})
	}
}
