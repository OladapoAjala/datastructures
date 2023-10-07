package doublehashing

import (
	"fmt"
	"hash/fnv"

	"github.com/OladapoAjala/datastructures/helpers"
	"github.com/OladapoAjala/datastructures/sets"
	"github.com/OladapoAjala/datastructures/sets/data"
	"golang.org/x/exp/constraints"
)

type HashTable[K constraints.Ordered] struct {
	Table      []*data.Data[K, any]
	capacity   int32
	size       int32
	loadFactor float32
}

const (
	MAX_LOAD_FACTOR  = 0.70
	DEFAULT_CAPACITY = 7
)

type HashTabler[K constraints.Ordered] interface {
	sets.Seter[K, any]
}

func NewHashTable[K constraints.Ordered](capacity int32) *HashTable[K] {
	var cap int32
	if capacity == 0 {
		cap = DEFAULT_CAPACITY
	} else if helpers.IsPrime(capacity) {
		cap = capacity
	} else {
		cap = helpers.NextPrime(capacity)
	}
	table := make([]*data.Data[K, any], cap)

	return &HashTable[K]{
		capacity:   cap,
		size:       0,
		loadFactor: 0,
		Table:      table,
	}
}

func (h *HashTable[K]) Insert(key K, value any) error {
	if key == *new(K) {
		return fmt.Errorf("invalid key")
	}
	item := data.NewDataWithHash(key, value)
	index := item.GetHash() % uint32(h.capacity)

	switch v := h.Table[index]; {
	case v == nil || v.IsTombStone():
		h.Table[index] = item
	case v.Key == key:
		v.Value = value
		return nil
	default:
		var x uint32 = 1
		h1 := index
		for v != nil {
			index = h.getNextIndex(v, h1, x)
			v = h.Table[index]
			x++
		}
		h.Table[index] = item
	}

	h.size++
	h.loadFactor = float32(h.size) / float32(h.capacity)
	if h.loadFactor >= MAX_LOAD_FACTOR {
		return h.resize()
	}
	return nil
}

func (h *HashTable[K]) resize() error {
	cap := helpers.NextPrime(h.GetCapacity())
	ht := NewHashTable[K](cap)

	for _, it := range h.Table {
		if it == nil || it.IsTombStone() {
			continue
		}
		err := ht.Insert(it.GetKey(), it.GetValue())
		if err != nil {
			return err
		}
	}

	h.Table = ht.Table
	h.capacity = ht.capacity
	h.size = ht.size
	h.loadFactor = ht.loadFactor
	return nil
}

func (h *HashTable[K]) Find(key K) (any, error) {
	index, err := h.getIndex(key)
	if err != nil {
		return nil, err
	}
	item := h.Table[index]
	return item.GetValue(), nil
}

func (h *HashTable[K]) getIndex(key K) (int32, error) {
	if key == *new(K) {
		return -1, fmt.Errorf("invalid key")
	}

	data, index := h.GetData(key)
	if data == nil {
		return -1, fmt.Errorf("key %v not found in hashtable", key)
	} else if data.Key == key {
		return int32(index), nil
	} else {
		var x uint32 = 1
		h1 := index

		for !data.IsTombStone() {
			index = h.getNextIndex(data, h1, x)
			data = h.Table[index]
			if data == nil {
				return -1, fmt.Errorf("key %v not found in hashtable", key)
			}
			if data.Key == key {
				return int32(index), nil
			}
			x++
		}
	}

	return -1, fmt.Errorf("key %v not found in hashtable", key)
}

func (h *HashTable[K]) Delete(key K) error {
	index, err := h.getIndex(key)
	if err != nil {
		return err
	}
	h.Table[index] = data.NewTombStone[K, any]()
	h.size--
	h.loadFactor = float32(h.size) / float32(h.capacity)
	return nil
}

func (h *HashTable[K]) GetData(key K) (*data.Data[K, any], uint32) {
	hasher := fnv.New32()
	hasher.Write([]byte(data.ToString(key)))
	index := hasher.Sum32() % uint32(h.capacity)
	return h.Table[index], index
}

func (h *HashTable[K]) GetCapacity() int32 {
	return h.capacity
}

func (h *HashTable[K]) GetSize() int32 {
	return h.size
}

func (h *HashTable[K]) GetLoadFactor() float32 {
	return h.loadFactor
}

func (h *HashTable[K]) getNextIndex(item *data.Data[K, any], offset, x uint32) uint32 {
	delta := item.Probe() % uint32(h.capacity)
	if delta == 0 {
		delta = 1
	}
	return (offset + (x * delta)) % uint32(h.capacity)
}
