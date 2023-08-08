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
	return da.array[index], nil
}

func (da *DynamicArray[T]) Contains(data T) bool {
	for _, d := range da.array {
		if d == data {
			return true
		}
	}
	return false
}

func (da *DynamicArray[T]) Insert(index int32, data T) error {
	if index >= da.capacity {
		newArray := make([]T, 2*da.capacity)
		for i, d := range da.array {
			newArray[i] = d
		}
		da.array = newArray
	}

	da.array[index] = data
	da.length += 1
	da.capacity = int32(len(da.array))
	return nil
}

func (da *DynamicArray[T]) InsertFirst(data T) error {
	return da.Insert(0, data)
}

func (da *DynamicArray[T]) InsertLast(data T) error {
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
	return da.Delete(da.length - 1)
}

func (da *DynamicArray[T]) Size() int32 {
	return da.length
}

func (da *DynamicArray[T]) Capacity() int32 {
	return da.capacity
}

// TODO: complete sort function (merge sort).
func (da *DynamicArray[T]) Sort() error {
	return nil
}

func (da *DynamicArray[T]) IsEmpty() bool {
	return da.length == 0
}

// TODO: complete ToLinkedList function.
func (da *DynamicArray[T]) ToLinkedList() (*linkedlist.LinkedList[T], error) {
	// sort.Sort(data)
	return nil, nil
}
