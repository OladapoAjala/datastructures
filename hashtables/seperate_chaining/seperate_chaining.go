package seperatechaining

import (
	"fmt"
	"hash/fnv"

	"github.com/OladapoAjala/datastructures/hashtables"
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
	"github.com/OladapoAjala/datastructures/sets/data"
	"golang.org/x/exp/constraints"
)

type HashTable[K constraints.Ordered] struct {
	capacity  int32
	size      int32
	threshold int32
	Table     []*linkedlist.LinkedList[*data.Data[K, any]]
}

type HashTabler[K constraints.Ordered] interface {
	hashtables.HashTabler[K]
	GetThreshold() int32
}

const MAX_LOAD_FACTOR float32 = 0.80

var _ HashTabler[string] = new(HashTable[string])

func NewHashTable[K constraints.Ordered](capacity int32) *HashTable[K] {
	table := make([]*linkedlist.LinkedList[*data.Data[K, any]], capacity)
	threshold := float32(capacity) * MAX_LOAD_FACTOR

	return &HashTable[K]{
		capacity:  capacity,
		size:      0,
		threshold: int32(threshold),
		Table:     table,
	}
}

func (h *HashTable[K]) Insert(key K, value any) error {
	if key == *new(K) {
		return fmt.Errorf("invalid key")
	}
	item := data.NewDataWithHash(key, value)
	pos := item.GetHash() % uint32(h.capacity)

	if isPresent, prev := h.contains(item, pos); isPresent {
		prev.Value = value
		return nil
	}

	if h.Table[pos] == nil {
		h.Table[pos] = linkedlist.NewList[*data.Data[K, any]]()
	}

	err := h.Table[pos].InsertLast(item)
	if err != nil {
		return err
	}
	h.size++

	if h.size > h.threshold {
		return h.resize()
	}
	return nil
}

func (h *HashTable[K]) resize() error {
	capacity := h.capacity * 2
	ht := NewHashTable[K](capacity)

	for _, ll := range h.Table {
		if ll == nil {
			continue
		}

		for i := int32(0); i < ll.GetSize(); i++ {
			node, err := ll.GetNode(i)
			if err != nil {
				return fmt.Errorf("error resizing table %w", err)
			}

			err = ht.Insert(node.Data.GetKey(), node.Data.GetValue())
			if err != nil {
				return fmt.Errorf("error resizing table %w", err)
			}
		}
	}

	h.capacity = ht.capacity
	h.size = ht.size
	h.threshold = ht.threshold
	h.Table = ht.Table
	return nil
}

func (h *HashTable[K]) contains(input *data.Data[K, any], pos uint32) (bool, *data.Data[K, any]) {
	ll := h.Table[pos]
	if ll == nil {
		return false, nil
	}

	for it := ll.Head; it != nil; it = it.Next {
		if input.Equal(it.Data) {
			return true, it.Data
		}
	}
	return false, nil
}

func (h *HashTable[K]) Find(key K) (any, error) {
	if key == *new(K) {
		return nil, fmt.Errorf("invalid key")
	}
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

	h.size--
	if h.Table[pos].GetSize() == 0 {
		h.Table[pos] = nil
	}

	return nil
}

func (h *HashTable[K]) getIndex(key K, ll *linkedlist.LinkedList[*data.Data[K, any]]) (int32, error) {
	if ll.IsEmpty() {
		return -1, fmt.Errorf("empty list")
	}

	data := data.NewDataWithHash[K, any](key, nil)
	var index int32 = 0
	for it := ll.Head; it != nil; it = it.Next {
		if data.Equal(it.Data) {
			return index, nil
		}
		index++
	}

	return -1, fmt.Errorf("key not found")
}

func (h *HashTable[K]) GetSize() int32 {
	return h.size
}

func (h *HashTable[K]) GetCapacity() int32 {
	return h.capacity
}

func (h *HashTable[K]) GetThreshold() int32 {
	return h.threshold
}
