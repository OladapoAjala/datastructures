package node

type Node[T any] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}

func NewNode[T any]() *Node[T] {
	return new(Node[T])
}
