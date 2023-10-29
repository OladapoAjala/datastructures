package trees

import (
	"github.com/OladapoAjala/datastructures/trees/node"
)

type ITrees[T comparable] interface {
	Successor(*node.Node[T]) (*node.Node[T], error)
	Predecessor(*node.Node[T]) (*node.Node[T], error)
	SubTreeFirst(*node.Node[T]) (*node.Node[T], error)
	SubTreeLast(*node.Node[T]) (*node.Node[T], error)
	TraversalOrder(*node.Node[T]) ([]T, error)
	PreOrderTraversal(*node.Node[T]) ([]T, error)
	PostOrderTraversal(*node.Node[T]) ([]T, error)
	GetSize() int32
	GetHeight() int32
	// SubtreeAugmentation()
	// GetNodeHeight(*node.Node[T]) int32
	// GetNodeDepth(*node.Node[T]) int32
	// Sequence() *sequences.Sequencer[T]
	// Set() *sets.Seter[T, any]
}
