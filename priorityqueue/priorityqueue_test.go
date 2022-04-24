package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dequeue(t *testing.T) {
	is := assert.New(t)
	pq := NewPQueue[string]()
	pq.Enqueue("A", "B", "D")

	tests := []struct {
		name   string
		pQueue *PQueue[string]
		want   func(string, error)
	}{
		{
			name:   "add element to the queue",
			pQueue: pq,
			want: func(data string, err error) {
				is.Nil(err)
				is.Equal(data, "A")
				is.Equal(pq.Tree, []string{"B", "D"})
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
