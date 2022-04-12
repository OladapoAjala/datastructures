package linkedlist

import (
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
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
			tt.list.Add(tt.args.data)
			if !tt.want(tt.list) {
				t.Errorf("append failed")
			}
		})
	}
}

func TestLinkedList_AddFirst(t *testing.T) {
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
			tt.list.AddFirst(tt.args.data)
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

func TestLinkedList_Remove(t *testing.T) {
	type args struct {
		index int32
	}
	tests := []struct {
		name string
		args args
		list *LinkedList[string]
		want func(*LinkedList[string], any) bool
	}{
		{
			name: "remove head",
			args: args{
				index: 0,
			},
			list: NewList("Node 0", "Node 1"),
			want: func(l *LinkedList[string], got any) bool {
				return l.Head.Next == nil && l.Head.Data == "Node 1" && got == "Node 0"
			},
		},
		{
			name: "remove tail",
			args: args{
				index: 2,
			},
			list: NewList("A", "B", "C"),
			want: func(l *LinkedList[string], got any) bool {
				return l.Tail.Data == "B" && got == "C"
			},
		},
		{
			name: "remove from list with three nodes",
			args: args{
				index: 1,
			},
			list: NewList("Node 0", "Node 1", "Node 2"),
			want: func(l *LinkedList[string], got any) bool {
				return l.Head.Next.Data == "Node 2" && l.Tail.Data == "Node 2" && got == "Node 1"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.list.Remove(tt.args.index)
			if err != nil {
				t.Errorf("error removing node from list, %v", err)
			}

			if !tt.want(tt.list, got) {
				t.Errorf("Error removing Node from list")
			}
		})
	}
}
