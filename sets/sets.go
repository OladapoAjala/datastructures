package sets

import (
	"github.com/OladapoAjala/datastructures/sets/data"
	"golang.org/x/exp/constraints"
)

type Seter[K constraints.Ordered, V any] interface {
	Find(K) V
	Insert(*data.Data[K, V]) error
	Delete(K) V
	InOrder() (*data.Data[K, V], error)
	FindMin() V
	FindMax() V
}
