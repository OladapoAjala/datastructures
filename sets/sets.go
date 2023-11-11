package sets

import (
	"golang.org/x/exp/constraints"
)

type Seter[K constraints.Ordered, V any] interface {
	Find(K) (V, error)
	Insert(K, V) error
	Delete(K) (V, error)
	FindMin() (V, error)
	FindMax() (V, error)
	FindNext(K) (V, error)
	FindPrev(K) (V, error)
	Size() int32
}
