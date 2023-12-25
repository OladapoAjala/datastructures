package priorityqueue

import (
	"github.com/OladapoAjala/datastructures/heap/maxheap"
	"golang.org/x/exp/constraints"
)

type PQueue[K constraints.Ordered, V comparable] struct {
	*maxheap.MaxHeap[K, V]
}

type IPQueue[T comparable] interface {
	Dequeue() (T, error)
	Enqueue(T) error
}

func NewPQueue[K constraints.Ordered, V comparable]() *PQueue[K, V] {
	return &PQueue[K, V]{
		maxheap.NewMaxHeap[K, V](),
	}
}

func (pq *PQueue[K, V]) Dequeue() (V, error) {
	max, err := pq.DeleteMax()
	if err != nil {
		return *new(V), err
	}
	return max.GetValue(), nil
}

func (pq *PQueue[K, V]) Enqueue(key K, val V) error {
	err := pq.Insert(key, val)
	if err != nil {
		return err
	}
	return nil
}
