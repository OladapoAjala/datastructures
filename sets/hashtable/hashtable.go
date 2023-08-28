package hashtable

import (
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
	"golang.org/x/exp/constraints"
)

type Entry[K constraints.Ordered] struct {
	hash  int
	key   K
	value any
}

func NewEntry[K constraints.Ordered](key K, value any) *Entry[K] {
	entry := &Entry[K]{
		key:   key,
		value: value,
	}
	// keyBytes := []byte(key)
	// entry.hash = utils.FNVHash(key)
	return entry
}

func (e *Entry[K]) equal(val *Entry[K]) bool {
	if e.hash != val.hash {
		return false
	}
	return e.key == val.key
}

type HashTable[K constraints.Ordered, V Entry[K]] struct {
	size  int
	Table []*linkedlist.LinkedList[string]
}

type IHashTable[K constraints.Ordered, V any] interface {
	Insert(K, V) error
	LookUp(K) (V, error)
	Remove(K) error
}

var _ IHashTable[string, any] = new(HashTable[string, Entry[string]])

func NewHashTable[K constraints.Ordered, V any]() *HashTable[K, Entry[K]] {
	return nil
}

func (h *HashTable[K, V]) Insert(key K, value any) error {
	// entry := NewEntry[K](key, value)
	// pos := entry.hash % h.size
	// h.Table[pos] = linkedlist.NewList[*V](entry)
	return nil
}

func (h *HashTable[K, V]) LookUp(key K) (any, error) {
	/*
		1. compute hash with key
		2. compute table index with hash
		3. go to table index and look for linkedlist node with key
		4. return value
	*/
	return nil, nil
}

func (h *HashTable[K, V]) Remove(key K) error {
	return nil
}
