package heap

import (
	"golang.org/x/exp/constraints"
)

type Heaper[K constraints.Ordered, V comparable] interface {
	Insert(K, V) error
}
