package sets

import (
	"github.com/OladapoAjala/datastructures/sets/data"
	"golang.org/x/exp/constraints"
)

type Seter[K constraints.Ordered, V any] interface {
	Find(K) (V, error)
	Insert(K, V) error
	Delete(K) (V, error)
	InOrder() ([]*data.Data[K, V], error)
	FindMin() (V, error)
	FindMax() (V, error)
	FindNext(K) (V, error)
	FindPrev(K) (V, error)
	Size() int32
}
