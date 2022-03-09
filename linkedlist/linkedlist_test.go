package linkedlist

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	node_0 := &node{
		index: 0,
		value: "Node 0",
		prev:  nil,
		next:  nil,
	}

	node_1 := &node{
		index: 1,
		value: "Node 1",
		prev:  node_0,
		next:  nil,
	}

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
				return l.head.value == "DSA sucks"

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
				return l.tail.value == "Node 2"
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
				return l.head.value == "DSA Sucks"
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
