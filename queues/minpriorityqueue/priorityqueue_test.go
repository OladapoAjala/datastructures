package minpriorityqueue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dequeue(t *testing.T) {
	is := assert.New(t)
	var pq *PQueue[int, string]

	tests := []struct {
		name   string
		setup  func() *PQueue[int, string]
		pQueue *PQueue[int, string]
		want   func(string, error)
	}{
		{
			name: "deque from queue with single element",
			setup: func() *PQueue[int, string] {
				pq := NewPQueue[int, string]()
				err := pq.Enqueue(0, "A")
				is.Nil(err)
				return pq
			},
			want: func(data string, err error) {
				is.Nil(err)
				is.Equal(data, "A")
				min, err := pq.FindMin()
				is.Error(err, fmt.Errorf("empty heap"))
				is.Nil(min)
			},
		},
		{
			name: "deque element",
			setup: func() *PQueue[int, string] {
				pq := NewPQueue[int, string]()
				is.Nil(pq.Enqueue(2, "C"))
				is.Nil(pq.Enqueue(0, "A"))
				is.Nil(pq.Enqueue(1, "B"))
				return pq
			},
			want: func(data string, err error) {
				is.Nil(err)
				is.Equal(data, "A")
				min, err := pq.FindMin()
				is.Nil(err)
				is.Equal(min.GetValue(), "B")
				is.EqualValues(pq.Heap.GetSize(), 2)
			},
		},
		{
			name: "deque element -- again",
			want: func(data string, err error) {
				is.Nil(err)
				is.Equal(data, "B")
				min, err := pq.FindMin()
				is.Nil(err)
				is.Equal(min.GetValue(), "C")
				is.EqualValues(pq.Heap.GetSize(), 1)
			},
		},
		{
			name: "deque element -- empty the queue",
			want: func(data string, err error) {
				is.Nil(err)
				is.Equal(data, "C")
				is.EqualValues(pq.Heap.GetSize(), 0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				pq = tt.setup()
			}
			data, err := pq.Dequeue()
			tt.want(data, err)
		})
	}
}
