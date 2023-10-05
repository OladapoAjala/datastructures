package doublehashing

import (
	"fmt"
	"hash/fnv"
	"math"

	"github.com/OladapoAjala/datastructures/sets"
	"github.com/OladapoAjala/datastructures/sets/data"
	"golang.org/x/exp/constraints"
)

var TOMBSTONE = data.Data[int, any]{
	Key:   -10928623050,
	Value: nil,
}

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
	} else if isPrime(capacity) {
		cap = capacity
	} else {
		cap = nextPrime(capacity)
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
	case v == nil:
		h.Table[index] = item
	case v.Key == key:
		v.Value = value
		return nil
	default:
		var x uint32 = 1
		h1 := index
		for h.Table[index] != nil {
			delta := item.Probe() % uint32(h.capacity)
			if delta == 0 {
				delta = 1
			}
			index = (h1 + (x * delta)) % uint32(h.capacity)
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
	cap := nextPrime(h.GetCapacity())
	ht := NewHashTable[K](cap)

	for _, it := range h.Table {
		if it == nil {
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
	if key == *new(K) {
		return nil, fmt.Errorf("invalid key")
	}
	hasher := fnv.New32()
	hasher.Write([]byte(data.ToString(key)))
	index := hasher.Sum32() % uint32(h.capacity)
	item := h.Table[index]

	if item == nil {
		return nil, fmt.Errorf("key %v not found in hashtable", key)
	}
	if item.Key != key {
		// compute new index
	}

	return item.Value, nil
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

func isPrime(num int32) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%int32(i) == 0 {
			return false
		}
	}
	return true
}

func nextPrime(input int32) int32 {
	input++
	for !isPrime(input) {
		input++
	}
	return input
}
