package sets

import (
	"github.com/OladapoAjala/datastructures/sets/data"
	"golang.org/x/exp/constraints"
)

type Seter[K constraints.Ordered, V any] interface {
	Find(K) V
	Insert(K, V) error
	Delete(K) (V, error)
	InOrder() (*data.Data[K, V], error)
	FindMin() V
	FindMax() V
	FindNext(K) V
	FindPrev(K) V
	Size() int32
}
