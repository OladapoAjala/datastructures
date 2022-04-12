package heap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	Tree []T
	Size int32
	Map  map[T][]int32 // you'd have to make sure the indices are sorted.
}

type IHeap[T constraints.Ordered] interface {
	Add(T) error
	Insert(T) error
	IsEmpty() bool
	Poll() (T, error)
	Remove(T) error
	sink(int32) error
	swap(int32, int32) error
	swim(int32) error
	mapAdd(T, int32) error
	mapSet(T, int32) error
	less(int32, int32) bool
}

// var _ IHeap[int] = new(Heap[int])

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

		err = h.swim(h.Size - 1)
		if err != nil {
			return err
		}

		err = h.mapAdd(d, h.Size)
		if err != nil {
			return err
		}

		h.Size += 1
	}

	return nil
}

func (h *Heap[T]) mapAdd(data T, index int32) error {
	var zero T
	if data == zero {
		return fmt.Errorf("cannot use null value as map key")
	}

	h.Map[data] = append(h.Map[data], index)
	// sort map indices in ascending order.
	return nil
}

func (h *Heap[T]) mapSet(data T, index int32) error {
	if _, ok := h.Map[data]; !ok {
		return fmt.Errorf("cannot update map data (%v) absent in tree %v", data, h.Tree)
	}

	for _, i := range h.Map[data] {
		if i == index {
			return nil
		}
	}

	h.Map[data] = append(h.Map[data], index)
	// sort map indices in ascending order.
	return nil
}

func (h *Heap[T]) swim(i int32) error {
	parenIndex := (i - 1) / 2

	var err error
	// maybe i or parenIndex
	for parenIndex > 0 && h.less(parenIndex, i) {
		err = h.swap(parenIndex, i)
		if err != nil {
			return err
		}

		i = parenIndex
		parenIndex = (i - 1) / 2
	}

	return err
}

func (h *Heap[T]) swap(i, j int32) error {
	h.Tree[i], h.Tree[j] = h.Tree[j], h.Tree[i]

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

func (h *Heap[T]) less(i, j int32) bool {
	return h.Tree[i] < h.Tree[j]
}
