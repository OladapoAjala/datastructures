package node

/*
Abstract Data Types:
1. List
2. Queue
3. Stack
4. Map
5.

Possible implementations
1. List: LinkedList, Dynamic Array
2. Queue: LinkedList, Dynamic Array
3. Stack: LinkedList, Dynamic Array
4. Map: Trees, Hash Map, Hash Table

*/

type Node struct {
	Data any
	Prev *Node
	Next *Node
}

func NewNode() *Node {
	return new(Node)
}
