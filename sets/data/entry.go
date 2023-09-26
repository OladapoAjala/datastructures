package data

import (
	"fmt"
	"hash/fnv"

	"golang.org/x/exp/constraints"
)

type Entry[K constraints.Ordered, V any] struct {
	*Data[K, V]
	hash uint32
}

func NewEntry[K constraints.Ordered, V any](key K, value V) *Entry[K, V] {
	hasher := fnv.New32()
	hasher.Write([]byte(fmt.Sprintf("%s", key)))

	return &Entry[K, V]{
		Data: NewData[K, V](key, value),
		hash: hasher.Sum32(),
	}
}

func (e *Entry[K, V]) Equal(val *Entry[K, V]) bool {
	if e.hash != val.hash {
		return false
	}
	return e.Key == val.Key
}

func (e *Entry[K, V]) GetHash() uint32 {
	return e.hash
}
