package trees

import (
	"github.com/OladapoAjala/datastructures/trees/data"
	"github.com/OladapoAjala/datastructures/trees/node"
	"golang.org/x/exp/constraints"
)

// TODO: migrate ITRees to ITree.

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
}

type ITree[K constraints.Ordered, V comparable] interface {
	Successor(*data.Data[K, V]) (*data.Data[K, V], error)
	Predecessor(*data.Data[K, V]) (*data.Data[K, V], error)
	SubTreeFirst(*data.Data[K, V]) (*data.Data[K, V], error)
	SubTreeLast(*data.Data[K, V]) (*data.Data[K, V], error)
	TraversalOrder(*data.Data[K, V]) ([]V, error)
	PreOrderTraversal(*data.Data[K, V]) ([]V, error)
	PostOrderTraversal(*data.Data[K, V]) ([]V, error)
	GetSize() int32
	GetHeight() int32
	// SubtreeAugmentation()
}
