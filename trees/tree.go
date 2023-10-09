package trees

import (
	"github.com/OladapoAjala/datastructures/sequences"
	"github.com/OladapoAjala/datastructures/trees/node"
)

type ITrees[T any] interface {
	Insert(T) error
	Remove(T) error

	Successor(*node.Node[T]) *node.Node[T] //immediate element after current element in in-order traversal (there are two cases)
	Predecessor(*node.Node[T]) *node.Node[T]

	SubTreeFirst(*node.Node[T])                       //leftmost node
	SubTreeInsertAfter(*node.Node[T], *node.Node[T])  // replace successor with new node
	SubTreeInsertBefore(*node.Node[T], *node.Node[T]) // replace predecessor with new node

	Sequence() *sequences.Sequencer[T]
	// Set() *sets.Seter[T, any]

	// Subtree(*node.Node[T], int32) (*node.Node[T], error)
	SubtreeAugmentation()
	GetNodeHeight(T) int32
	GetNodeDepth(T) int32

	// Order
	TraversalOrder(*node.Node[T]) ([]T, error)
}
