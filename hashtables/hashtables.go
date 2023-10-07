package hashtables

import (
	"golang.org/x/exp/constraints"
)

type HashTabler[K constraints.Ordered] interface {
	Insert(K, any) error
	Find(K) (any, error)
	Delete(K) error
	GetSize() int32
	GetCapacity() int32
}
