package queue

import (
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name  string
		queue *Queue[string]
		args  args
		want  func(*Queue[string]) bool
	}{
		{
			name:  "enqueue element",
			queue: NewQueue[string](),
			args: args{
				data: "A",
			},
			want: func(q *Queue[string]) bool {
				return q.Head.Data == "A" && q.Length == 1
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.queue.Enqueue(tt.args.data)
			if err != nil {
				t.Errorf("Queue.Peek() error = %v", err)
				return
			}

			if !tt.want(tt.queue) {
				t.Errorf("Queue.Enqueue() error = %v", err)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	testQueue := NewQueue[string]()
	testQueue.Enqueue("A")
	testQueue.Enqueue("B")
	testQueue.Enqueue("C")

	tests := []struct {
		name  string
		queue *Queue[string]
		want  func(*Queue[string], any) bool
	}{
		{
			name:  "dequeue element",
			queue: testQueue,
			want: func(q *Queue[string], got any) bool {
				return q.Head.Data == "B" && got == "A" && q.Length == 2
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.queue.Dequeue()
			if err != nil {
				t.Errorf("Stack.Peek() error = %v", err)
				return
			}

			if !tt.want(tt.queue, got) {
				t.Errorf("Queue.Enqueue() error, got = %v", got)
			}
		})
	}
}
