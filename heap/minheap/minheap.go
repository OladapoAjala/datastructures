package minheap

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/heap"
	"github.com/OladapoAjala/datastructures/heap/data"
	"github.com/OladapoAjala/datastructures/sequences/dynamicarray"
	"golang.org/x/exp/constraints"
)

type MinHeap[K constraints.Ordered, V comparable] struct {
	Heap *dynamicarray.DynamicArray[*data.Data[K, V]]
}

type MinHeaper[K constraints.Ordered, V comparable] interface {
	heap.Heaper[K, V]
	FindMin() (*data.Data[K, V], error)
	DeleteMin() (*data.Data[K, V], error)
}

var _ MinHeaper[string, string] = new(MinHeap[string, string])

func NewMinHeap[K constraints.Ordered, V comparable]() *MinHeap[K, V] {
	return &MinHeap[K, V]{
		Heap: dynamicarray.NewDynamicArray[*data.Data[K, V]](),
	}
}

func (mh *MinHeap[K, V]) Insert(key K, value V) error {
	index := mh.Heap.GetSize()
	d := data.NewData[K, V](key, value, index)
	err := mh.Heap.InsertLast(d)
	if err != nil {
		return err
	}
	return mh.heapifyUp(d)
}

func (mh *MinHeap[K, V]) FindMin() (*data.Data[K, V], error) {
	if mh.IsEmpty() {
		return nil, fmt.Errorf("empty heap")
	}
	return mh.Heap.GetData(0)
}

func (mh *MinHeap[K, V]) DeleteMin() (*data.Data[K, V], error) {
	if mh.IsEmpty() {
		return nil, fmt.Errorf("empty heap")
	}

	min, err := mh.Heap.GetData(0)
	if err != nil {
		return nil, err
	}
	last, err := mh.Heap.GetData(mh.Heap.GetSize() - 1)
	if err != nil {
		return nil, err
	}
	mh.swap(min, last)
	err = mh.Heap.DeleteLast()
	if err != nil {
		return nil, err
	}
	err = mh.heapifyDown(min)
	if err != nil {
		return nil, err
	}
	return last, nil
}

func (mh *MinHeap[K, V]) heapifyUp(d *data.Data[K, V]) error {
	if d == nil || d.Index == 0 {
		return nil
	}
	parent, err := mh.GetParent(d)
	if err != nil {
		return err
	}
	if d.GetKey() < parent.GetKey() {
		mh.swap(d, parent)
	}
	return mh.heapifyUp(parent)
}

func (mh *MinHeap[K, V]) heapifyDown(d *data.Data[K, V]) error {
	if d == nil {
		return nil
	}

	left, err := mh.GetLeft(d)
	if err != nil {
		return nil
	}
	right, err := mh.GetRight(d)
	if err != nil {
		return nil
	}

	if right.GetKey() < left.GetKey() {
		mh.swap(d, right)
		return mh.heapifyDown(right)
	}
	mh.swap(d, left)
	return mh.heapifyDown(left)
}

func (mh *MinHeap[K, V]) swap(a, b *data.Data[K, V]) {
	a.Key, b.Key = b.Key, a.Key
	a.Value, b.Value = b.Value, a.Value
}

func (mh *MinHeap[K, V]) GetParent(n *data.Data[K, V]) (*data.Data[K, V], error) {
	parentIndex := n.GetParentIndex()
	return mh.Heap.GetData(parentIndex)
}

func (mh *MinHeap[K, V]) GetLeft(n *data.Data[K, V]) (*data.Data[K, V], error) {
	leftIndex := n.GetLeftIndex()
	return mh.Heap.GetData(leftIndex)
}

func (mh *MinHeap[K, V]) GetRight(n *data.Data[K, V]) (*data.Data[K, V], error) {
	rightIndex := n.GetRightIndex()
	return mh.Heap.GetData(rightIndex)
}

func (mh *MinHeap[K, V]) IsEmpty() bool {
	return mh.Heap.GetSize() == 0
}
