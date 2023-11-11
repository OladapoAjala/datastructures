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
	InOrder() ([]*data.Data[K, V], error)
}

var _ ISortedArray[int, string] = new(SortedArray[int, string])

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
	copy(sa.array, sort[K, V](values))
	return sa
}

func sort[K constraints.Ordered, V any](input []*data.Data[K, V]) []*data.Data[K, V] {
	if len(input) <= 2 {
		if len(input) == 2 {
			if input[0].Key > input[1].Key {
				input[0], input[1] = input[1], input[0]
			}
		}
		return input
	}

	mid := len(input) / 2
	arr1 := sort(input[:mid])
	arr2 := sort(input[mid:])

	return merge(arr1, arr2)
}

func merge[K constraints.Ordered, V any](arrOne, arrTwo []*data.Data[K, V]) []*data.Data[K, V] {
	p1, p2 := 0, 0
	sorted := make([]*data.Data[K, V], 0)
	for p1 < len(arrOne) && p2 < len(arrTwo) {
		if arrOne[p1].Key <= arrTwo[p2].Key {
			sorted = append(sorted, arrOne[p1])
			p1++
			continue
		}

		sorted = append(sorted, arrTwo[p2])
		p2++
	}

	sorted = append(sorted, arrOne[p1:]...)
	sorted = append(sorted, arrTwo[p2:]...)
	return sorted
}

func (sa *SortedArray[K, V]) Insert(key K, value V) error {
	if key == *new(K) {
		return fmt.Errorf("empty key")
	}

	index, err := sa.getIndex(key)
	if err == nil {
		sa.array[index] = data.NewData(key, value)
		return nil
	}

	item := data.NewData(key, value)
	sa.array[sa.lenght] = item
	sa.lenght++
	sort(sa.array[:sa.lenght])

	if sa.lenght == sa.capacity {
		newArr := make([]*data.Data[K, V], 2*sa.capacity)
		copy(newArr, sa.array)
		sa.array = newArr
		sa.capacity = int32(len(newArr))
	}
	return nil
}

func (sa *SortedArray[K, V]) Delete(key K) (V, error) {
	index, err := sa.getIndex(key)
	if err != nil {
		return *new(V), err
	}

	output := sa.array[index]
	sa.array[index] = nil
	sa.shift(index)
	sa.lenght--
	return output.Value, nil
}

func (sa *SortedArray[K, V]) getIndex(key K) (int32, error) {
	if sa.GetLenght() == 0 {
		return -1, fmt.Errorf("empty array")
	}

	min := int32(0)
	max := sa.GetLenght() - 1

	for min <= max {
		mid := (min + max) / 2

		if sa.array[mid].GetKey() == key {
			return mid, nil
		}

		if sa.array[mid].GetKey() > key {
			max = mid - 1
			continue
		}

		min = mid + 1
	}

	return -1, fmt.Errorf("key: %v not found", key)
}

func (sa *SortedArray[K, V]) shift(index int32) {
	for i := index; i < sa.GetLenght(); i++ {
		sa.array[i] = sa.array[i+1]
	}
}

func (sa *SortedArray[K, V]) Find(key K) (V, error) {
	index, err := sa.getIndex(key)
	if err != nil {
		return *new(V), err
	}
	return sa.array[index].GetValue(), nil
}

func (sa *SortedArray[K, V]) FindMin() (V, error) {
	if sa.IsEmpty() {
		return *new(V), fmt.Errorf("empty array")
	}
	return sa.array[0].GetValue(), nil
}

func (sa *SortedArray[K, V]) FindMax() (V, error) {
	if sa.IsEmpty() {
		return *new(V), fmt.Errorf("empty array")
	}
	return sa.array[sa.GetLenght()-1].GetValue(), nil
}

func (sa *SortedArray[K, V]) FindNext(key K) (V, error) {
	index, err := sa.getIndex(key)
	if err != nil {
		return *new(V), err
	}
	return sa.array[index+1].Value, nil
}

func (sa *SortedArray[K, V]) FindPrev(key K) (V, error) {
	index, err := sa.getIndex(key)
	if err != nil {
		return *new(V), err
	}
	return sa.array[index-1].Value, nil
}

func (sa *SortedArray[K, V]) InOrder() ([]*data.Data[K, V], error) {
	return sa.array, nil
}

func (sa *SortedArray[K, V]) GetLenght() int32 {
	return sa.lenght
}

func (sa *SortedArray[K, V]) GetCapacity() int32 {
	return sa.capacity
}

func (sa *SortedArray[K, V]) IsEmpty() bool {
	return sa.lenght == 0
}

func (sa *SortedArray[K, V]) IsSorted() bool {
	for i := int32(0); i < sa.GetLenght()-1; i++ {
		if sa.array[i].Key > sa.array[i+1].Key {
			return false
		}
	}
	return true
}

func (sa *SortedArray[K, V]) Size() int32 {
	return sa.GetLenght()
}
