package queue

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/linkedlist"
)

type Queue struct {
	*linkedlist.LinkedList
}

type IQueue interface {
	Dequeue() (any, error)
	Enqueue(data any) error
}

var _ IQueue = new(Queue)

func NewQueue() *Queue {
	return &Queue{
		new(linkedlist.LinkedList),
	}
}

func (q *Queue) Enqueue(data any) error {
	err := q.Add(data)
	if err != nil {
		return fmt.Errorf("error enqueuing data %v, error: %v", data, err)
	}
	return nil
}

func (q *Queue) Dequeue() (any, error) {
	data, err := q.Remove(0)
	if err != nil {
		return nil, fmt.Errorf("unable to remove first element, %v", err)
	}

	return data, nil
}
