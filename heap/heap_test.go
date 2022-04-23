package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	is := assert.New(t)
	h := NewHeap[int]()

	type args struct {
		data []int
	}

	tests := []struct {
		name string
		args args
		heap *Heap[int]
		want func(*Heap[int], error)
	}{
		{
			name: "add [3,2,1] to heap",
			args: args{
				data: []int{3, 2, 1},
			},
			heap: h,
			want: func(h *Heap[int], err error) {
				is.Nil(err)
				is.Equal(h.Tree, []int{1, 3, 2})
				is.Equal(h.Map, map[int][]int32{1: {0}, 3: {1}, 2: {2}})
			},
		},
		{
			name: "add [7,6,3,4,1,5,2] to heap",
			args: args{
				data: []int{7, 6, 3, 4, 1, 5, 2},
			},
			heap: NewHeap[int](),
			want: func(h *Heap[int], err error) {
				is.Nil(err)
				is.Equal(h.Tree, []int{1, 3, 2, 7, 4, 6, 5})
				is.Equal(h.Map, map[int][]int32{1: {0}, 3: {1}, 2: {2}, 7: {3}, 4: {4}, 6: {5}, 5: {6}})
			},
		},
		{
			name: "add [7,2,3,4,1,3,2] to heap",
			args: args{
				data: []int{7, 2, 3, 4, 1, 3, 2},
			},
			heap: NewHeap[int](),
			want: func(h *Heap[int], err error) {
				is.Nil(err)
				is.Equal(h.Tree, []int{1, 2, 2, 7, 4, 3, 3})
				is.Equal(h.Map, map[int][]int32{1: {0}, 2: {1, 2}, 3: {5, 6}, 4: {4}, 7: {3}})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.heap.Add(tt.args.data...)
			tt.want(tt.heap, err)
		})
	}
}

func Test_Remove(t *testing.T) {
	h := NewHeap[int]()
	h.Add(7, 2, 3, 4, 1, 3, 2)

	type args struct {
		data int
	}

	tests := []struct {
		name string
		args args
		heap *Heap[int]
		// want func(*Heap[int], error)
	}{
		{
			name: "remove element & f*ing sink",
			args: args{
				data: 7,
			},
			heap: h,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.heap.Remove(tt.args.data)
			if err != nil {
				t.Error("unexpected error")
			}
		})
	}

}

func Test_less(t *testing.T) {
	h := NewHeap[string]()
	h.Add("A", "B", "C")

	type args struct {
		i, j int32
	}

	tests := []struct {
		name string
		args args
		heap *Heap[string]
	}{
		{
			name: "compare strings",
			args: args{
				i: 0,
				j: 1,
			},
			heap: h,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.heap.less(tt.args.i, tt.args.j) {
				t.Errorf("expected A to be lesser than B")
			}
		})
	}
}

func Test_mapAdd(t *testing.T) {
	is := assert.New(t)
	h := NewHeap[string]()
	h.Add("A", "B", "C")

	type args struct {
		data string
	}

	tests := []struct {
		name string
		args args
		heap *Heap[string]
		want func(*Heap[string], error)
	}{
		{
			name: "add D to map",
			args: args{
				data: "D",
			},
			heap: h,
			want: func(h *Heap[string], err error) {
				is.Nil(err)
				is.Equal(h.Map, map[string][]int32{
					"A": {0},
					"B": {1},
					"C": {2},
					"D": {3},
				})
			},
		},
		{
			name: "add already existing key",
			args: args{
				data: "A",
			},
			heap: h,
			want: func(h *Heap[string], err error) {
				is.Nil(err)
				is.Equal(h.Map, map[string][]int32{
					"A": {0, 3},
					"B": {1},
					"C": {2},
					"D": {3},
				})
			},
		},
		{
			name: "add empty data",
			args: args{
				data: "",
			},
			heap: h,
			want: func(h *Heap[string], err error) {
				is.Error(err, "cannot use null value as map key")
				is.Equal(h.Map, map[string][]int32{
					"A": {0, 3},
					"B": {1},
					"C": {2},
					"D": {3},
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.heap.mapAdd(tt.args.data)
			tt.want(tt.heap, err)
		})
	}
}

func Test_mapSet(t *testing.T) {
	is := assert.New(t)
	h := NewHeap[string]()
	h.Add("A", "B", "C")

	type args struct {
		data  string
		index int32
	}

	tests := []struct {
		name string
		args args
		heap *Heap[string]
		want func(*Heap[string], error)
	}{
		{
			name: "simple test case",
			args: args{
				data:  "C",
				index: 5,
			},
			heap: h,
			want: func(h *Heap[string], err error) {
				is.Nil(err)
				is.Equal(h.Map, map[string][]int32{
					"A": {0},
					"B": {1},
					"C": {2, 5},
				})
			},
		},
		{
			name: "set an already existing index",
			args: args{
				data:  "B",
				index: 1,
			},
			heap: h,
			want: func(h *Heap[string], err error) {
				is.Error(err, "data (B) already at index (1)")
				is.Equal(h.Map, map[string][]int32{
					"A": {0},
					"B": {1},
					"C": {2, 5},
				})
			},
		},
		{
			name: "set index for key absent in map",
			args: args{
				data:  "D",
				index: 3,
			},
			heap: h,
			want: func(h *Heap[string], err error) {
				is.Nil(err)
				is.Equal(h.Map, map[string][]int32{
					"A": {0},
					"B": {1},
					"C": {2, 5},
					"D": {3},
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.heap.mapSet(tt.args.data, tt.args.index)
			tt.want(tt.heap, err)
		})
	}
}

func Test_swim(t *testing.T) {
	is := assert.New(t)
	h := NewHeap[string]()
	h.Add("B", "C", "A")

	type args struct {
		i int32
	}

	tests := []struct {
		name string
		args args
		heap *Heap[string]
		want func(*Heap[string], error)
	}{
		{
			name: "swim simple structure",
			args: args{
				i: 2,
			},
			heap: h,
			want: func(h *Heap[string], err error) {
				is.Nil(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.heap.swim(tt.args.i)
			tt.want(tt.heap, err)
		})
	}
}
