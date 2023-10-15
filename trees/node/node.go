package node

type Node[T any] struct {
	Data   T
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (n *Node[T]) GetData() T {
	return n.Data
}

func (n *Node[T]) GetParent() *Node[T] {
	return n.Parent
}

func (n *Node[T]) GetLeft() *Node[T] {
	return n.Left
}

func (n *Node[T]) GetRight() *Node[T] {
	return n.Right
}

func (n *Node[T]) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}
