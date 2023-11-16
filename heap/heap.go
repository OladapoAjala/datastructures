package heap

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type Heaper[K constraints.Ordered, V comparable] interface {
	Insert(K, V) error
}
type Heap[T constraints.Ordered] struct {
	Tree []T
	Size int32
	Map  map[T][]int32
}

type IHeap[T constraints.Ordered] interface {
	Add(...T) error
	IsEmpty() bool
	Poll() (T, error)
	Remove(T) error
	less(int32, int32) bool
	mapAdd(T) error
	mapSet(T, int32) error
	mapSwap(int32, int32) error
	removeMapIndex(T, int32)
	sink(int32) error
	// sortMap(T)
	swap(int32, int32) error
	swim(int32) error
}

var _ IHeap[int] = new(Heap[int])

func NewHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		Tree: make([]T, 0),
		Size: 0,
		Map:  make(map[T][]int32, 0),
	}
}

func (h *Heap[T]) Add(data ...T) error {
	if len(data) < 1 {
		return fmt.Errorf("cannot add null value to heap.")
	}

	var err error
	for _, d := range data {
		h.Tree = append(h.Tree, d)
		err = h.mapAdd(d)
		if err != nil {
			return err
		}
		h.Size += 1
		// DEPRECATE this in favour of the heapify method O(n).
		// err = h.swim(h.Size - 1)
		// if err != nil {
		// 	return err
		// }
	}

	var i int32
	for i = (h.Size / 2) - 1; i >= 0; i-- {
		h.sink(i)
	}
	return nil
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Size == 0
}

func (h *Heap[T]) Poll() (T, error) {
	var zero T
	if h.IsEmpty() {
		return zero, fmt.Errorf("empty heap")
	}

	data := h.Tree[0]
	err := h.Remove(data)
	if err != nil {
		return zero, err
	}
	return data, nil
}

func (h *Heap[T]) Remove(data T) error {
	var zero T
	if data == zero {
		return fmt.Errorf("cannot use null value as map key")
	}

	if _, ok := h.Map[data]; !ok {
		return fmt.Errorf("value is absent in heap")
	}

	index := h.Map[data][0]
	h.swap(index, h.Size-1)
	h.Tree = h.Tree[:h.Size-1]
	h.removeMapIndex(data, h.Size-1)
	h.Size--

	if index == int32(math.Max(0, float64(h.Size-1))) {
		return nil
	}

	currElem := h.Tree[index]
	h.sink(index)

	if h.Tree[index] == currElem {
		h.swim(index)
	}
	return nil
}

func (h *Heap[T]) less(i, j int32) bool {
	return h.Tree[i] < h.Tree[j]
}

func (h *Heap[T]) mapAdd(data T) error {
	var zero T
	if data == zero {
		return fmt.Errorf("cannot use null value as map key")
	}

	h.Map[data] = append(h.Map[data], h.Size)
	// h.sortMap(data)
	return nil
}

func (h *Heap[T]) mapSet(data T, index int32) error {
	if _, ok := h.Map[data]; !ok {
		h.Map[data] = make([]int32, 0)
	}

	for _, i := range h.Map[data] {
		if i == index {
			return fmt.Errorf("data (%v) already at index (%d)", data, index)
		}
	}

	h.Map[data] = append(h.Map[data], index)
	// h.sortMap(data)
	return nil
}

func (h *Heap[T]) mapSwap(i, j int32) error {
	// TODO: [1,2,2] what happens when i,j = 1,2
	h.removeMapIndex(h.Tree[i], i)
	h.removeMapIndex(h.Tree[j], j)

	err := h.mapSet(h.Tree[i], j)
	if err != nil {
		return err
	}

	h.mapSet(h.Tree[j], i)
	if err != nil {
		return err
	}

	return nil
}

func (h *Heap[T]) removeMapIndex(key T, i int32) {
	oldIndices := h.Map[key]
	newIndices := make([]int32, 0)

	for _, e := range oldIndices {
		if e == i {
			continue
		}
		newIndices = append(newIndices, e)
	}
	h.Map[key] = newIndices

	if len(h.Map[key]) == 0 {
		delete(h.Map, key)
	}
}

func (h *Heap[T]) sink(k int32) error {
	if k >= h.Size {
		return fmt.Errorf("index out of range")
	}

	for {
		leftNode := 2*k + 1
		rightNode := 2*k + 2
		smallest := leftNode

		if rightNode < h.Size && h.less(rightNode, leftNode) {
			smallest = rightNode
		}
		if leftNode >= h.Size || h.less(k, smallest) {
			break
		}

		err := h.swap(smallest, k)
		if err != nil {
			return err
		}
		k = smallest
	}

	return nil
}

/*
-> I'm guessing the indices are sorted already

	func (h *Heap[T]) sortMap(data T) {
		value := h.Map[data]
		sort.Slice(value, func(i, j int) bool {
			return value[i] < value[j]
		})
		h.Map[data] = value
	}
*/
func (h *Heap[T]) swap(i, j int32) error {
	err := h.mapSwap(i, j)
	if err != nil {
		return err
	}

	h.Tree[i], h.Tree[j] = h.Tree[j], h.Tree[i]
	return nil
}

func (h *Heap[T]) swim(i int32) error {
	parentIndex := (i - 1) / 2

	var err error
	for parentIndex >= 0 && h.less(i, parentIndex) {
		err = h.swap(i, parentIndex)
		if err != nil {
			return err
		}

		i = parentIndex
		parentIndex = (i - 1) / 2
	}

	return err
}
