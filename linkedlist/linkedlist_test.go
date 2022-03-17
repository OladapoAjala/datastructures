package linkedlist

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	testList := NewList("Node 0", "Node 1", "Node 2")

	type args struct {
		data interface{}
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
				data: "DSA sucks",
			},
			want: func(l *LinkedList) bool {
				return l.Head.Data == "DSA sucks" && l.Tail.Data == "DSA sucks"

			},
		},
		{
			name: "append to a list with three Nodes",
			list: testList,
			args: args{
				data: "Node 3",
			},
			want: func(l *LinkedList) bool {
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
		data interface{}
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
				data: "DSA Sucks",
			},
			want: func(l *LinkedList) bool {
				return l.Head.Data == "DSA Sucks" && l.Tail.Data == "DSA Sucks"
			},
		},
		{
			name: "prepend to a list with three Nodes",
			list: testList,
			args: args{
				data: "Node 3",
			},
			want: func(l *LinkedList) bool {
				return l.Head.Data == "Node 3" && l.Tail.Data == "Node 2" && l.Tail.Index == 3 && l.Head.Index == 0
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

// func TestLinkedList_Insert(t *testing.T) {
// 	Node_0.next = Node_1

// 	type args struct {
// 		index int32
// 		value string
// 	}
// 	tests := []struct {
// 		name string
// 		list *LinkedList
// 		args args
// 		want func(l *LinkedList) bool
// 	}{
// 		{
// 			name: "insert into a list with two Nodes",
// 			list: &LinkedList{
// 				length: 2,
// 				head:   Node_0,
// 				tail:   Node_1,
// 			},
// 			args: args{
// 				index: 1,
// 				value: "Node 3",
// 			},
// 			want: func(l *LinkedList) bool {
// 				return l.head.next.value == "Node 3"
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.list.Insert(tt.args.index, tt.args.value)

// 			if !tt.want(tt.list) {
// 				t.Errorf("Error inserting into list")
// 			}
// 		})
// 	}
// }

// func TestLinkedList_Remove(t *testing.T) {
// 	Node_0.next = Node_1

// 	type args struct {
// 		index int32
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		list *LinkedList
// 		want func(l *LinkedList) bool
// 	}{
// 		{
// 			name: "remove single Node from list with two Nodes",
// 			args: args{
// 				index: 1,
// 			},
// 			list: &LinkedList{
// 				length: 2,
// 				head:   Node_0,
// 				tail:   Node_1,
// 			},
// 			want: func(l *LinkedList) bool {
// 				return l.head.next == nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.list.Remove(tt.args.index)

// 			if !tt.want(tt.list) {
// 				t.Errorf("Error removing Node from list")
// 			}
// 		})
// 	}
// }

// func TestLinkedList_ShiftLeft(t *testing.T) {
// 	type args struct {
// 		index int32
// 	}

// 	tests := []struct {
// 		name string
// 		args args
// 		list *LinkedList
// 		want func(l *LinkedList) bool
// 	}{
// 		{
// 			name: "simple node shift",
// 			args: args{
// 				index: 1,
// 			},
// 			list: NewList(),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.list.ShiftLeft(tt.args.index)

// 			if !tt.want(tt.list) {
// 				t.Errorf("error shifting node to the left")
// 			}
// 		})
// 	}
// }

// func TestLinkedList_ShiftRight(t *testing.T) {
// 	type fields struct {
// 		Length int32
// 		Head   *node.Node
// 		Tail   *node.Node
// 	}
// 	type args struct {
// 		index int32
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   func(l *LinkedList)
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				Length: tt.fields.Length,
// 				Head:   tt.fields.Head,
// 				Tail:   tt.fields.Tail,
// 			}
// 			if err := l.ShiftRight(tt.args.index); (err != nil) != tt.wantErr {
// 				t.Errorf("LinkedList.ShiftRight() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
