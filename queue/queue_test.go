package queue

import (
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	type args struct {
		data any
	}
	tests := []struct {
		name  string
		queue *Queue
		args  args
		want  func(*Queue) bool
	}{
		{
			name:  "enqueue element",
			queue: NewQueue(),
			args: args{
				data: "A",
			},
			want: func(q *Queue) bool {
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
	testQueue := NewQueue()
	testQueue.Enqueue("A")
	testQueue.Enqueue("B")
	testQueue.Enqueue("C")

	tests := []struct {
		name  string
		queue *Queue
		want  func(*Queue, any) bool
	}{
		{
			name:  "dequeue element",
			queue: testQueue,
			want: func(q *Queue, got any) bool {
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
