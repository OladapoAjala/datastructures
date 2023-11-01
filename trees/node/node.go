package node

type Node[T any] struct {
	Data   T
	Size   int32
	Height int32
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

type Noder[T any] interface {
	GetData() T
	GetParent() *Node[T]
	GetLeft() *Node[T]
	GetRight() *Node[T]
	GetHeight() int32
	GetSize() int32
	IsLeaf() bool
	Skew() int32
}

var _ Noder[any] = new(Node[any])

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		Data:  data,
		Size:  1,
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

func (n *Node[T]) GetHeight() int32 {
	return n.Height
}

func (n *Node[T]) GetSize() int32 {
	return n.Size
}

func (n *Node[T]) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node[T]) Skew() int32 {
	var hr, hl int32
	if n.Right != nil {
		hr = n.Right.Height
	}
	if n.Left != nil {
		hl = n.Left.Height
	}
	return hr - hl
}

func (n *Node[T]) RightRotate() {
	parent := n.Parent
	left := n.Left
	n.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = n
	}
	left.Right = n
	n.Parent = left
	left.Parent = parent
	if parent != nil {
		if parent.Left == n {
			parent.Left = left
		} else {
			parent.Right = left
		}
	}
}

func (n *Node[T]) LeftRotate() {
	parent := n.Parent
	right := n.Right
	n.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = n
	}
	right.Left = n
	n.Parent = right
	right.Parent = parent
	if parent != nil {
		if parent.Right == n {
			parent.Right = right
		} else {
			parent.Left = right
		}
	}
}
