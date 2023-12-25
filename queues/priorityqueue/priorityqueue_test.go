package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dequeue(t *testing.T) {
	is := assert.New(t)
	pq := NewPQueue[int, string]()
	is.Nil(pq.Enqueue(-2, "C"))
	is.Nil(pq.Enqueue(0, "A"))
	is.Nil(pq.Enqueue(-1, "B"))

	tests := []struct {
		name   string
		pQueue *PQueue[int, string]
		want   func(string, error)
	}{
		{
			name:   "deque element",
			pQueue: pq,
			want: func(data string, err error) {
				is.Nil(err)
				is.Equal(data, "A")
				max, err := pq.FindMax()
				is.Nil(err)
				is.Equal(max.GetValue(), "B")
				is.EqualValues(pq.Heap.GetSize(), 2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.pQueue.Dequeue()
			tt.want(data, err)
		})
	}
}
