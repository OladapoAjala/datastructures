package heap

import (
	"testing"
)

func Test_mapSet(t *testing.T) {

	h := NewHeap[int]()
	h.Add(1, 2, 3)

	type args struct {
		data  int
		index int32
	}

	tests := []struct {
		name string
		args args
		heap *Heap[int]
		want func(*Heap[int]) bool
	}{
		{
			name: "simple test case",
			args: args{
				data:  3,
				index: 5,
			},
			heap: h,
			want: func(h *Heap[int]) bool {
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.mapSet(tt.args.data, tt.args.index)

			if !tt.want(tt.heap) {
				t.Errorf("map set failed")
			}
		})
	}
}
