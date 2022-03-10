package linkedlist

import (
	"testing"
)

var (
	node_0 = &node{
		index: 0,
		value: "Node 0",
		prev:  nil,
		next:  nil,
	}

	node_1 = &node{
		index: 1,
		value: "Node 1",
		prev:  node_0,
		next:  nil,
	}
)

func TestLinkedList_Append(t *testing.T) {
	node_0.next = node_1

	type args struct {
		value string
	}
	tests := []struct {
		name string
		list *LinkedList
		args args
		want func(l *LinkedList) bool
	}{
		{
			name: "append to an empty list",
			list: new(LinkedList),
			args: args{
				value: "DSA sucks",
			},
			want: func(l *LinkedList) bool {
				return l.head.value == "DSA sucks" && l.tail.value == "DSA sucks"

			},
		},
		{
			name: "append to a list with two nodes",
			list: &LinkedList{
				length: 2,
				head:   node_0,
				tail:   node_1,
			},
			args: args{
				value: "Node 2",
			},
			want: func(l *LinkedList) bool {
				return l.tail.value == "Node 2" && l.head.value == "Node 0"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Append(tt.args.value)
			if !tt.want(tt.list) {
				t.Errorf("append failed")
			}
		})
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	node_0.next = node_1

	type args struct {
		value string
	}
	tests := []struct {
		name string
		list *LinkedList
		args args
		want func(l *LinkedList) bool
	}{
		{
			name: "prepend to an empty list",
			list: new(LinkedList),
			args: args{
				value: "DSA Sucks",
			},
			want: func(l *LinkedList) bool {
				return l.head.value == "DSA Sucks" && l.tail.value == "DSA Sucks"
			},
		},
		{
			name: "prepend to a list with two nodes",
			list: &LinkedList{
				length: 2,
				head:   node_0,
				tail:   node_1,
			},
			args: args{
				value: "Node 3",
			},
			want: func(l *LinkedList) bool {
				return l.head.value == "Node 3" && l.tail.value == "Node 1"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Prepend(tt.args.value)
			if !tt.want(tt.list) {
				t.Errorf("prepend failed")
			}
		})
	}
}

func TestLinkedList_Insert(t *testing.T) {
	node_0.next = node_1

	type args struct {
		index int32
		value string
	}
	tests := []struct {
		name string
		list *LinkedList
		args args
		want func(l *LinkedList) bool
	}{
		{
			name: "insert into a list with two nodes",
			list: &LinkedList{
				length: 2,
				head:   node_0,
				tail:   node_1,
			},
			args: args{
				index: 1,
				value: "Node 3",
			},
			want: func(l *LinkedList) bool {
				return l.head.next.value == "Node 3"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Insert(tt.args.index, tt.args.value)

			if !tt.want(tt.list) {
				t.Errorf("Error inserting into list")
			}
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	node_0.next = node_1

	type args struct {
		index int32
	}
	tests := []struct {
		name string
		args args
		list *LinkedList
		want func(l *LinkedList) bool
	}{
		{
			name: "remove single node from list with two nodes",
			args: args{
				index: 1,
			},
			list: &LinkedList{
				length: 2,
				head:   node_0,
				tail:   node_1,
			},
			want: func(l *LinkedList) bool {
				return l.head.next == nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Remove(tt.args.index)

			if !tt.want(tt.list) {
				t.Errorf("Error removing node from list")
			}
		})
	}
}
