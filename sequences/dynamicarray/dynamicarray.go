package dynamicarray

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/sequences"
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
)

type DynamicArray[T comparable] struct {
	array    []T
	length   int32
	capacity int32
}

type IDynamicArray[T comparable] interface {
	sequences.Sequencer[T]
	ToLinkedList() (*linkedlist.LinkedList[T], error)
	Capacity() int32
}

var _ IDynamicArray[string] = new(DynamicArray[string])

func NewDynamicArray[T comparable](data ...T) *DynamicArray[T] {
	da := new(DynamicArray[T])

	if len(data) == 0 {
		da.array = make([]T, 1)
		da.length = 0
		da.capacity = 1
		return da
	}

	da.array = make([]T, 2*len(data))
	da.length = int32(len(data))
	da.capacity = int32(len(da.array))
	copy(da.array, data)
	return da
}

func (da *DynamicArray[T]) GetData(index int32) (T, error) {
	if index >= da.capacity {
		return *new(T), fmt.Errorf("index out of range")
	}
	if index >= da.length {
		return *new(T), fmt.Errorf("index out of range")
	}
	return da.array[index], nil
}

func (da *DynamicArray[T]) Contains(data T) bool {
	for i, d := range da.array {
		if i == int(da.length) {
			return false
		}
		if d == data {
			return true
		}
	}
	return false
}

func (da *DynamicArray[T]) Insert(index int32, data T) error {
	if index >= da.capacity {
		newArray := make([]T, 2*da.capacity)
		copy(newArray, da.array)
		da.array = newArray
		da.capacity = int32(len(da.array))
		return da.Insert(index, data)
	}

	da.array[index] = data
	da.capacity = int32(len(da.array))
	if index >= da.length {
		da.length = index + 1
	}
	return nil
}

func (da *DynamicArray[T]) InsertFirst(data T) error {
	return da.Insert(0, data)
}

func (da *DynamicArray[T]) InsertLast(data T) error {
	if da.length == 0 {
		return da.Insert(0, data)
	}
	return da.Insert(da.length-1, data)
}

func (da *DynamicArray[T]) Delete(index int32) error {
	if index >= da.length {
		return fmt.Errorf("index out of range")
	}

	da.array[index] = *new(T)
	return nil
}

func (da *DynamicArray[T]) DeleteFirst() error {
	return da.Delete(0)
}

func (da *DynamicArray[T]) DeleteLast() error {
	if da.length == 0 {
		return da.Delete(0)
	}
	return da.Delete(da.length - 1)
}

func (da *DynamicArray[T]) Size() int32 {
	return da.length
}

func (da *DynamicArray[T]) Capacity() int32 {
	return da.capacity
}

func (da *DynamicArray[T]) IsEmpty() bool {
	return da.length == 0
}

// TODO: complete ToLinkedList function.
func (da *DynamicArray[T]) ToLinkedList() (*linkedlist.LinkedList[T], error) {
	// sort.Sort(data)
	return nil, nil
}
