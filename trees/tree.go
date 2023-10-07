package trees

import (
	"github.com/OladapoAjala/datastructures/sequences"
	"github.com/OladapoAjala/datastructures/sets"
	"github.com/OladapoAjala/datastructures/trees/node"
	"golang.org/x/exp/constraints"
)

type ITrees[T constraints.Ordered] interface {
	SubTreeFirst(*node.Node[T]) //leftmost node
	Successor(*node.Node[T])    //immediate element after current element in in-order traversal (there are two cases)
	Predecessor(*node.Node[T])
	SubTreeInsertAfter(*node.Node[T], *node.Node[T])  // replace successor with new node
	SubTreeInsertBefore(*node.Node[T], *node.Node[T]) // replace predecessor with new node
	Sequence() *sequences.Sequencer[T]
	Set() *sets.Seter[T, any]
	Subtree(*node.Node[T], int32)
	SubtreeAugmentation()
}
