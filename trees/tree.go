package trees

import (
	"github.com/OladapoAjala/datastructures/trees/node"
)

type ITrees[T comparable] interface {
	InsertAfter(*node.Node[T], T) error
	InsertBefore(*node.Node[T], T) error
	Successor(*node.Node[T]) (*node.Node[T], error)
	Predecessor(*node.Node[T]) (*node.Node[T], error)
	SubTreeFirst(*node.Node[T]) (*node.Node[T], error)
	SubTreeLast(*node.Node[T]) (*node.Node[T], error)
	TraversalOrder(*node.Node[T]) ([]T, error)
	PreOrderTraversal(*node.Node[T]) ([]T, error)
	PostOrderTraversal(*node.Node[T]) ([]T, error)

	// Subtree(*node.Node[T], int32) (*node.Node[T], error)
	// SubtreeAugmentation()
	// GetNodeHeight(*node.Node[T]) int32
	// GetNodeDepth(*node.Node[T]) int32

	// Delete(T) error
	// Insert(T) error
	// Sequence() *sequences.Sequencer[T]
	// Set() *sets.Seter[T, any]

	GetSize() int32
	GetHeight() int32
}
