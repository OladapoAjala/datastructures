package queue

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/linkedlist"
)

type Queue[T any] struct {
	*linkedlist.LinkedList[T]
}

type IQueue[T any] interface {
	Dequeue() (T, error)
	Enqueue(T) error
}

var _ IQueue[int] = new(Queue[int])

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		new(linkedlist.LinkedList[T]),
	}
}

func (q *Queue[T]) Enqueue(data T) error {
	err := q.Add(data)
	if err != nil {
		return fmt.Errorf("error enqueuing data %v, error: %v", data, err)
	}
	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	data, err := q.Remove(0)
	if err != nil {
		return zero, fmt.Errorf("unable to remove first element, %v", err)
	}

	return data, nil
}
