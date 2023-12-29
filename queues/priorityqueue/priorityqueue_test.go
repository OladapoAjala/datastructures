package priorityqueue

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
		want   func(int, string, error)
	}{
		{
			name: "deque from queue with single element",
			setup: func() *PQueue[int, string] {
				pq := NewPQueue[int, string]()
				err := pq.Enqueue(0, "A")
				is.Nil(err)
				return pq
			},
			want: func(key int, value string, err error) {
				is.Nil(err)
				is.Equal(value, "A")
				is.Equal(key, 0)
				max, err := pq.FindMax()
				is.Error(err, fmt.Errorf("empty heap"))
				is.Nil(max)
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
			want: func(key int, value string, err error) {
				is.Nil(err)
				is.Equal(value, "C")
				is.Equal(key, 2)
				max, err := pq.FindMax()
				is.Nil(err)
				is.Equal(max.GetValue(), "B")
				is.EqualValues(pq.Heap.GetSize(), 2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq = tt.setup()
			key, value, err := pq.Dequeue()
			tt.want(key, value, err)
		})
	}
}
