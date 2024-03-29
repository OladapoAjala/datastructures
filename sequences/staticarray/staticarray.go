package staticarray

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/sequences"
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
)

type StaticArray[T comparable] struct {
	array  []T
	length int32
}

type IStaticArray[T comparable] interface {
	sequences.Sequencer[T]
	ToLinkedList() (*linkedlist.LinkedList[T], error)
}

var _ IStaticArray[string] = new(StaticArray[string])

func NewStaticArray[T comparable](size int32, data ...T) *StaticArray[T] {
	sa := new(StaticArray[T])
	sa.array = make([]T, size)
	copy(sa.array, data)
	sa.length = int32(len(sa.array))
	return sa
}

func (sa *StaticArray[T]) GetData(index int32) (T, error) {
	if index >= sa.length {
		return *new(T), fmt.Errorf("index out of range")
	}
	return sa.array[index], nil
}

func (sa *StaticArray[T]) Contains(data T) bool {
	for _, d := range sa.array {
		if d == data {
			return true
		}
	}
	return false
}

func (sa *StaticArray[T]) Set(index int32, data T) error {
	if index >= sa.length {
		return fmt.Errorf("index out of range")
	}

	sa.array[index] = data
	return nil
}

func (sa *StaticArray[T]) Insert(index int32, data T) error {
	if index >= sa.length {
		return fmt.Errorf("index out of range")
	}

	newArray := make([]T, sa.length)
	for i := int32(0); i < index; i++ {
		newArray[i] = sa.array[i]
	}
	newArray[index] = data
	for i := index; i < sa.length-1; i++ {
		newArray[i+1] = sa.array[i]
	}
	sa.array = newArray
	return nil
}

func (sa *StaticArray[T]) InsertFirst(data T) error {
	return sa.Insert(0, data)
}

func (sa *StaticArray[T]) InsertLast(data T) error {
	return sa.Insert(sa.length-1, data)
}

func (sa *StaticArray[T]) Delete(index int32) error {
	if index >= sa.length {
		return fmt.Errorf("index out of range")
	}

	sa.shift(index)
	return nil
}

func (sa *StaticArray[T]) shift(index int32) {
	for i := index; i < sa.GetSize()-1; i++ {
		sa.array[i] = sa.array[i+1]
	}
	sa.array[sa.GetSize()-1] = *new(T)
}

func (sa *StaticArray[T]) DeleteFirst() error {
	return sa.Delete(0)
}

func (sa *StaticArray[T]) DeleteLast() error {
	return sa.Delete(sa.length - 1)
}

func (sa *StaticArray[T]) GetSize() int32 {
	return sa.length
}

func (sa *StaticArray[T]) IsEmpty() bool {
	return sa.length == 0
}

// TODO: complete ToLinkedList function.
func (sa *StaticArray[T]) ToLinkedList() (*linkedlist.LinkedList[T], error) {
	return nil, nil
}
