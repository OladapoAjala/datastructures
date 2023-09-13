package sortedarray

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/sets"
	"github.com/OladapoAjala/datastructures/sets/data"
	"golang.org/x/exp/constraints"
)

type SortedArray[K constraints.Ordered, V any] struct {
	array    []*data.Data[K, V]
	lenght   int32
	capacity int32
}

type ISortedArray[K constraints.Ordered, V any] interface {
	sets.Seter[K, V]
}

// var _ ISortedArray[int, string] = new(SortedArray[int, string])

func NewSortedArray[K constraints.Ordered, V any](values ...*data.Data[K, V]) *SortedArray[K, V] {
	sa := new(SortedArray[K, V])

	if len(values) == 0 {
		sa.array = make([]*data.Data[K, V], 1)
		sa.lenght = 0
		sa.capacity = 1
		return sa
	}

	sa.array = make([]*data.Data[K, V], 2*len(values))
	sa.lenght = int32(len(values))
	sa.capacity = int32(len(sa.array))
	copy(sa.array, values)

	sort[K, V](sa)
	return sa
}

// Sort with mergesort.
func sort[K constraints.Ordered, V any](sa *SortedArray[K, V]) {
	if sa.Lenght() == 0 {
		return
	}

	sortHelper[K, V](sa.array, sa.Lenght()-1)
}

func sortHelper[K constraints.Ordered, V any](arr []*data.Data[K, V], lim int32) {
	if lim == 0 {
		return
	}

	l := getLargest[K, V](arr, lim)
	arr[lim], arr[l] = arr[l], arr[lim]
	sortHelper[K, V](arr, lim-1)
}

func getLargest[K constraints.Ordered, V any](arr []*data.Data[K, V], lim int32) int32 {
	if lim == 0 {
		return lim
	}

	val := getLargest[K, V](arr, lim-1)
	if arr[val].GetKey() > arr[lim].GetKey() {
		return val
	}
	return lim
}

func (sa *SortedArray[K, V]) Find(key K) (V, error) {
	if sa.Lenght() == 0 {
		return *new(V), fmt.Errorf("empty array")
	}

	min := int32(0)
	max := sa.Lenght() - 1

	for min <= max {
		mid := (min + max) / 2

		if sa.array[mid].GetKey() == key {
			return sa.array[mid].GetValue(), nil
		}

		if sa.array[mid].GetKey() > key {
			max = mid - 1
			continue
		}

		min = mid + 1
	}

	return *new(V), fmt.Errorf("key: %v not found", key)
}

func (sa *SortedArray[K, V]) FindMin(key K) V {
	return sa.array[0].GetValue()
}

func (sa *SortedArray[K, V]) FindMax() V {
	return sa.array[sa.Lenght()-1].GetValue()
}

func (sa *SortedArray[K, V]) Lenght() int32 {
	return sa.lenght
}

func (sa *SortedArray[K, V]) Capacity() int32 {
	return sa.capacity
}

func (da *SortedArray[K, V]) IsEmpty() bool {
	return da.lenght == 0
}
