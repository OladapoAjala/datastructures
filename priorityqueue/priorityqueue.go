package priorityqueue

import "github.com/OladapoAjala/datastructures/heap"

type PQueue[T comparable] struct {
	*heap.Heap[T]
}

type IPQueue[T comparable] interface {
	Dequeue() (T, error)
	Enqueue(T) error
}

func NewPQueue[T comparable](size int32) *PQueue[T] {
	return &PQueue[T]{
		heap.NewHeap[T](size),
	}
}

func NewPQueueWithElements[T comparable](elems ...T) *PQueue[T] {
	pq := NewPQueue[T](int32(len(elems)))

	for _, e := range elems {
		pq.Add(e)
	}

	// you want to sink all the elements to ensure they satisfy the heap invariant.

	return pq
}
