package maxheap

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/heap/data"
	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	is := assert.New(t)
	mh := NewMaxHeap[int, string]()

	tests := []struct {
		name string
		key  int
		val  string
		want func(error)
	}{
		{
			name: "Insert into an empty heap",
			key:  1,
			val:  "value1",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(1, mh.Heap.GetSize())

				max, err := mh.FindMax()
				is.Nil(err)
				is.Equal(max.Value, "value1")
				is.EqualValues(max.Index, 0)
			},
		},
		{
			name: "Insert at the end",
			key:  2,
			val:  "value2",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(2, mh.Heap.GetSize())

				max, err := mh.FindMax()
				is.Nil(err)
				is.Equal(max.Value, "value2")
				is.EqualValues(max.Index, 0)
			},
		},
		{
			name: "Insert at the beginning",
			key:  0,
			val:  "value0",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(3, mh.Heap.GetSize())

				max, err := mh.FindMax()
				is.Nil(err)
				is.Equal(max.Value, "value2")
				is.EqualValues(max.Index, 0)
			},
		},
		{
			name: "Insert with same key",
			key:  2,
			val:  "new_value",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(4, mh.Heap.GetSize())

				max, err := mh.FindMax()
				is.Nil(err)
				is.Equal(max.Value, "value2")
				is.EqualValues(max.Index, 0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mh.Insert(tt.key, tt.val)
			tt.want(err)
		})
	}
}

func Test_FindMax(t *testing.T) {
	is := assert.New(t)
	mh := NewMaxHeap[int, string]()

	tests := []struct {
		name string
		keys []int
		want func(*data.Data[int, string], error)
	}{
		{
			name: "FindMax in an empty heap",
			keys: []int{},
			want: func(max *data.Data[int, string], err error) {
				is.Error(err, fmt.Errorf("empty heap"))
				is.Nil(max)
			},
		},
		{
			name: "FindMax in a non-empty heap",
			keys: []int{3, 1, 5, 2, 4},
			want: func(max *data.Data[int, string], err error) {
				is.Nil(err)
				is.Equal(max.Value, "value5")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range tt.keys {
				mh.Insert(key, fmt.Sprintf("value%d", key))
			}
			max, err := mh.FindMax()
			tt.want(max, err)
		})
	}
}

func Test_DeleteMax(t *testing.T) {
	is := assert.New(t)
	mh := NewMaxHeap[int, string]()

	tests := []struct {
		name string
		keys []int
		want func(*data.Data[int, string], error)
	}{
		{
			name: "DeleteMax from an empty heap",
			keys: []int{},
			want: func(max *data.Data[int, string], err error) {
				is.Error(err, fmt.Errorf("empty heap"))
				is.Nil(max)
			},
		},
		{
			name: "DeleteMax from heap with one element",
			keys: []int{3},
			want: func(max *data.Data[int, string], err error) {
				is.Nil(err)
				is.Equal(max.Value, "value3")
			},
		},
		{
			name: "DeleteMax from a non-empty heap",
			keys: []int{3, 1, 5, 2, 4, 7},
			want: func(max *data.Data[int, string], err error) {
				is.Nil(err)
				is.Equal(max.Value, "value7")
			},
		},
		{
			name: "DeleteMax from a non-empty heap",
			want: func(max *data.Data[int, string], err error) {
				is.Nil(err)
				is.Equal(max.Value, "value5")
			},
		},
		{
			name: "DeleteMax from a non-empty heap",
			want: func(max *data.Data[int, string], err error) {
				is.Nil(err)
				is.Equal(max.Value, "value4")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range tt.keys {
				is.Nil(mh.Insert(key, fmt.Sprintf("value%d", key)))
			}
			max, err := mh.DeleteMax()
			tt.want(max, err)
		})
	}
}
