package minpriorityqueue

import (
	"github.com/OladapoAjala/datastructures/heap/minheap"
	"golang.org/x/exp/constraints"
)

type PQueue[K constraints.Ordered, V comparable] struct {
	*minheap.MinHeap[K, V]
}

type IPQueue[K constraints.Ordered, V comparable] interface {
	Dequeue() (V, error)
	Enqueue(K, V) error
}

func NewPQueue[K constraints.Ordered, V comparable]() *PQueue[K, V] {
	return &PQueue[K, V]{
		minheap.NewMinHeap[K, V](),
	}
}

func (pq *PQueue[K, V]) Dequeue() (K, V, error) {
	min, err := pq.DeleteMin()
	if err != nil {
		return *new(K), *new(V), err
	}
	return min.GetKey(), min.GetValue(), nil
}

func (pq *PQueue[K, V]) Enqueue(key K, val V) error {
	err := pq.Insert(key, val)
	if err != nil {
		return err
	}
	return nil
}
