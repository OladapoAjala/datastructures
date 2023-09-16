package queue

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
)

type Queue[T comparable] struct {
	*linkedlist.LinkedList[T]
}

type IQueue[T comparable] interface {
	Dequeue() (T, error)
	Enqueue(T) error
}

var _ IQueue[int] = new(Queue[int])

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		new(linkedlist.LinkedList[T]),
	}
}

func (q *Queue[T]) Enqueue(data T) error {
	err := q.InsertLast(data)
	if err != nil {
		return fmt.Errorf("error enqueuing data %v, error: %v", data, err)
	}
	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	data, err := q.GetData(0)
	if err != nil {
		return zero, fmt.Errorf("unable to remove first element, %v", err)
	}

	err = q.DeleteFirst()
	if err != nil {
		return zero, fmt.Errorf("unable to remove first element, %v", err)
	}

	return data, nil
}
