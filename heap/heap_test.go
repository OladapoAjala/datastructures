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
				is.Equal(h.Tree, []int{1, 2, 3})
				is.Equal(h.Map, map[int][]int32{1: {0}, 2: {1}, 3: {2}})
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
				is.Equal(h.Tree, []int{1, 4, 2, 7, 6, 5, 3})
				is.Equal(h.Map, map[int][]int32{1: {0}, 4: {1}, 2: {2}, 7: {3}, 6: {4}, 5: {5}, 3: {6}})
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
				is.Equal(h.Tree, []int{1, 2, 2, 4, 7, 3, 3})
				is.Equal(h.Map, map[int][]int32{1: {0}, 2: {2, 1}, 3: {5, 6}, 7: {4}, 4: {3}})
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

func Test_Poll(t *testing.T) {
	is := assert.New(t)
	h := NewHeap[string]()
	h.Add("D", "A", "C")

	data, err := h.Poll()
	is.Nil(err)
	is.Equal(data, "A")
	is.Equal(h.Tree, []string{"C", "D"})
	is.Equal(h.Map, map[string][]int32{"C": {0}, "D": {1}})

	data, err = h.Poll()
	is.Nil(err)
	is.Equal(data, "C")
	is.Equal(h.Tree, []string{"D"})
	is.Equal(h.Map, map[string][]int32{"D": {0}})

	data, err = h.Poll()
	is.Nil(err)
	is.Equal(data, "D")
	is.Equal(h.Tree, []string{})
	is.Equal(h.Map, map[string][]int32{})

	data, err = h.Poll()
	is.Empty(data)
	is.Error(err, "empty heap")
}

func Test_Remove(t *testing.T) {
	is := assert.New(t)
	h := NewHeap[int]()
	h.Add(7, 2, 3, 4, 1, 3, 2)

	h1 := NewHeap[int]()
	h1.Add(1, 7, 3, 8, 9, 5, 4)

	type args struct {
		data int
	}

	tests := []struct {
		name string
		args args
		heap *Heap[int]
		want func(*Heap[int], error)
	}{
		{
			name: "remove data from heap",
			args: args{
				data: 7,
			},
			heap: h,
			want: func(h *Heap[int], err error) {
				is.Nil(err)
				is.Equal(h.Tree, []int{1, 2, 2, 4, 3, 3})
				is.Equal(h.Map, map[int][]int32{1: {0}, 2: {2, 1}, 3: {5, 4}, 4: {3}})
			},
		},
		{
			name: "remove invalid data from heap",
			args: args{
				data: 10,
			},
			heap: h,
			want: func(h *Heap[int], err error) {
				is.Error(err, "value is absent in heap")
				is.Equal(h.Tree, []int{1, 2, 2, 4, 3, 3})
				is.Equal(h.Map, map[int][]int32{1: {0}, 2: {2, 1}, 3: {5, 4}, 4: {3}})
			},
		},
		{
			name: "test swim from remove",
			args: args{
				data: 8,
			},
			heap: h1,
			want: func(h *Heap[int], err error) {
				is.Nil(err)
				is.Equal(h.Tree, []int{1, 4, 3, 7, 9, 5})
				is.Equal(h.Map, map[int][]int32{1: {0}, 4: {1}, 3: {2}, 7: {3}, 9: {4}, 5: {5}})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.heap.Remove(tt.args.data)
			tt.want(tt.heap, err)
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
