package priorityqueue

import (
	"github.com/OladapoAjala/datastructures/heap/maxheap"
	"golang.org/x/exp/constraints"
)

type IQueuer[K any] interface {
	Insert(K)
	DeleteMax()
	FindMax()
}

type PQueue[T constraints.Ordered] struct {
	*maxheap.MaxHeap[T, T]
}

type IPQueue[T comparable] interface {
	Dequeue() (T, error)
	Enqueue(T) error
}

func NewPQueue[T constraints.Ordered]() *PQueue[T] {
	return &PQueue[T]{
		maxheap.NewMaxHeap[T, T](),
	}
}

func (pq *PQueue[T]) Dequeue() (T, error) {
	max, err := pq.DeleteMax()
	if err != nil {
		return *new(T), err
	}
	return max.GetKey(), nil
}

func (pq *PQueue[T]) Enqueue(data ...T) error {
	for _, d := range data {
		err := pq.Insert(d, d)
		if err != nil {
			return err
		}
	}
	return nil
}
