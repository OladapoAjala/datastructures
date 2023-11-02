package priorityqueue

import (
	"github.com/OladapoAjala/datastructures/heap"
	"golang.org/x/exp/constraints"
)

type IQueuer[K any] interface {
	Insert(K)
	DeleteMax()
	FindMax()
}

type PQueue[T constraints.Ordered] struct {
	*heap.Heap[T]
}

type IPQueue[T comparable] interface {
	Dequeue() (T, error)
	Enqueue(T) error
}

func NewPQueue[T constraints.Ordered]() *PQueue[T] {
	return &PQueue[T]{
		heap.NewHeap[T](),
	}
}

func (pq *PQueue[T]) Dequeue() (T, error) {
	return pq.Poll()
}

func (pq *PQueue[T]) Enqueue(data ...T) error {
	return pq.Add(data...)
}
