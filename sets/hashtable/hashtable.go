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
	capacity  int32
	size      int32
	threshold int32
	Table     []*linkedlist.LinkedList[*data.Entry[K]]
}

type HashTabler[K constraints.Ordered] interface {
	sets.Seter[K, any]
}

const MAX_LOAD_FACTOR float32 = 0.80

// var _ HashTabler[string, any] = new(HashTable[string, data.Entry[string]])

func NewHashTable[K constraints.Ordered](capacity int32) *HashTable[K] {
	table := make([]*linkedlist.LinkedList[*data.Entry[K]], capacity)
	threshold := float32(capacity) * MAX_LOAD_FACTOR

	return &HashTable[K]{
		capacity:  capacity,
		threshold: int32(threshold),
		Table:     table,
	}
}

func (h *HashTable[K]) Insert(key K, value any) error {
	if key == *new(K) {
		return fmt.Errorf("invalid key")
	}
	entry := data.NewEntry(key, value)
	pos := entry.GetHash() % uint32(h.capacity)

	if h.Table[pos] == nil {
		h.Table[pos] = linkedlist.NewList(entry)
		h.size++
		return nil
	}

	if isTrue, entry := h.contains(entry, pos); isTrue {
		entry.Value = value
		return nil
	}

	err := h.Table[pos].InsertLast(entry)
	if err != nil {
		return err
	}
	h.size++
	return nil
}

func (h *HashTable[K]) contains(entry *data.Entry[K], pos uint32) (bool, *data.Entry[K]) {
	ll := h.Table[pos]
	if ll == nil {
		return false, nil
	}

	for it := ll.Head; it != nil; it = it.Next {
		if entry.Equal(it.Data) {
			return true, it.Data
		}
	}
	return false, nil
}

func (h *HashTable[K]) Find(key K) (any, error) {
	// Check for invalid keys
	hasher := fnv.New32()
	hasher.Write([]byte(data.ToString(key)))
	pos := hasher.Sum32() % uint32(h.capacity)

	if h.Table[pos] == nil {
		return nil, fmt.Errorf("key %v not found in hashtable", key)
	}

	return h.Table[pos].Head.Data.GetValue(), nil
}

func (h *HashTable[K]) Delete(key K) error {
	hasher := fnv.New32()
	hasher.Write([]byte(data.ToString(key)))
	pos := hasher.Sum32() % uint32(h.capacity)

	if h.Table[pos] == nil {
		return fmt.Errorf("key %v not found in hashtable", key)
	}

	ll := h.Table[pos]
	index, err := h.getIndex(key, ll)
	if err != nil {
		return err
	}

	err = h.Table[pos].Delete(index)
	if err != nil {
		return fmt.Errorf("key %v not found in hashtable", key)
	}

	if h.Table[pos].Size() == 0 {
		h.Table[pos] = nil
	}

	return nil
}

func (h *HashTable[K]) getIndex(key K, ll *linkedlist.LinkedList[*data.Entry[K]]) (int32, error) {
	if ll.IsEmpty() {
		return -1, fmt.Errorf("empty list")
	}

	entry := data.NewEntry[K](key, nil)
	var index int32 = 0
	for it := ll.Head; it != nil; it = it.Next {
		if entry.Equal(it.Data) {
			return index, nil
		}
		index++
	}

	return -1, fmt.Errorf("key not found")
}

func (h *HashTable[K]) Size() int32 {
	return h.capacity
}
