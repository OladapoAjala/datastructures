package hashtable

import (
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
	"github.com/OladapoAjala/datastructures/sets"
	"github.com/OladapoAjala/datastructures/sets/data"
	"golang.org/x/exp/constraints"
)

type HashTable[K constraints.Ordered] struct {
	size  int32
	Table []*linkedlist.LinkedList[*data.Entry[K, any]]
}

type HashTabler[K constraints.Ordered, V any] interface {
	sets.Seter[K, V]
}

// var _ HashTabler[string, any] = new(HashTable[string, data.Entry[string]])

func NewHashTable[K constraints.Ordered]() *HashTable[K] {
	table := make([]*linkedlist.LinkedList[*data.Entry[K, any]], 10)
	return &HashTable[K]{
		size:  int32(len(table)),
		Table: table,
	}
}

func (h *HashTable[K]) Insert(key K, value any) error {
	entry := data.NewEntry(key, value)
	pos := entry.GetHash() % uint32(h.size)

	if h.Table[pos] != nil {
		return h.Table[pos].InsertLast(entry)
	}

	h.Table[pos] = linkedlist.NewList(entry)
	return nil
}

func (h *HashTable[K]) Find(key K) (any, error) {
	/*
		1. compute hash with key
		2. compute table index with hash
		3. go to table index and look for linkedlist node with key
		4. return value
	*/
	return nil, nil
}

func (h *HashTable[K]) Delete(key K) error {
	return nil
}

func (h *HashTable[K]) Size() int32 {
	return h.size
}
