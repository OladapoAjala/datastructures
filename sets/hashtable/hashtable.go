package hashtable

import (
	"fmt"
	"hash/fnv"

	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
	"github.com/OladapoAjala/datastructures/sets"
	"github.com/OladapoAjala/datastructures/sets/data"
	"golang.org/x/exp/constraints"
)

type HashTable[K constraints.Ordered] struct {
	size  int32
	Table []*linkedlist.LinkedList[*data.Entry[K]]
}

type HashTabler[K constraints.Ordered] interface {
	sets.Seter[K, any]
}

// var _ HashTabler[string, any] = new(HashTable[string, data.Entry[string]])

func NewHashTable[K constraints.Ordered](size int32) *HashTable[K] {
	table := make([]*linkedlist.LinkedList[*data.Entry[K]], size)
	return &HashTable[K]{
		size:  int32(len(table)),
		Table: table,
	}
}

func (h *HashTable[K]) Insert(key K, value any) error {
	entry := data.NewEntry(key, value)
	pos := entry.GetHash() % uint32(h.size)

	if h.Table[pos] != nil {
		if h.contains(entry, pos) {
			return fmt.Errorf("key: %v, value: %v already in hash table", key, value)
		}
		return h.Table[pos].InsertLast(entry)
	}

	h.Table[pos] = linkedlist.NewList(entry)
	return nil
}

func (h *HashTable[K]) contains(entry *data.Entry[K], pos uint32) bool {
	ll := h.Table[pos]
	if ll == nil {
		return false
	}

	for it := ll.Head; it != nil; it = it.Next {
		if entry.Equal(it.Data) {
			return true
		}
	}
	return false
}

func (h *HashTable[K]) Find(key K) (any, error) {
	hasher := fnv.New32()
	hasher.Write([]byte(data.ToString(key)))
	pos := hasher.Sum32() % uint32(h.size)

	if h.Table[pos] == nil {
		return nil, fmt.Errorf("key %v not found in hashtable", key)
	}

	return h.Table[pos].Head.Data.GetValue(), nil
}

func (h *HashTable[K]) Delete(key K) error {
	hasher := fnv.New32()
	hasher.Write([]byte(data.ToString(key)))
	pos := hasher.Sum32() % uint32(h.size)

	if h.Table[pos] == nil {
		return fmt.Errorf("key %v not found in hashtable", key)
	}

	err := h.Table[pos].DeleteFirst()
	if err != nil {
		return fmt.Errorf("key %v not found in hashtable", key)
	}

	if h.Table[pos].Size() == 0 {
		h.Table[pos] = nil
	}

	return nil
}

func (h *HashTable[K]) Size() int32 {
	return h.size
}
