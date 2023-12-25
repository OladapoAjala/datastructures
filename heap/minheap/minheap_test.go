package minheap

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/heap/data"
	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	is := assert.New(t)
	mh := NewMinHeap[int, string]()

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

				min, err := mh.FindMin()
				is.Nil(err)
				is.Equal(min.Value, "value1")
				is.EqualValues(min.Index, 0)
			},
		},
		{
			name: "Insert at the end",
			key:  2,
			val:  "value2",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(2, mh.Heap.GetSize())

				min, err := mh.FindMin()
				is.Nil(err)
				is.Equal(min.Value, "value1")
				is.EqualValues(min.Index, 0)
			},
		},
		{
			name: "Insert at the beginning",
			key:  0,
			val:  "value0",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(3, mh.Heap.GetSize())

				min, err := mh.FindMin()
				is.Nil(err)
				is.Equal(min.Value, "value0")
				is.EqualValues(min.Index, 0)
			},
		},
		{
			name: "Insert with same key",
			key:  2,
			val:  "new_value",
			want: func(err error) {
				is.Nil(err)
				is.EqualValues(4, mh.Heap.GetSize())

				min, err := mh.FindMin()
				is.Nil(err)
				is.Equal(min.Value, "value0")
				is.EqualValues(min.Index, 0)
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

func Test_FindMin(t *testing.T) {
	is := assert.New(t)
	mh := NewMinHeap[int, string]()

	tests := []struct {
		name string
		keys []int
		want func(*data.Data[int, string], error)
	}{
		{
			name: "FindMin in an empty heap",
			keys: []int{},
			want: func(min *data.Data[int, string], err error) {
				is.Error(err, fmt.Errorf("empty heap"))
				is.Nil(min)
			},
		},
		{
			name: "FindMin in a non-empty heap",
			keys: []int{3, 1, 5, 2, 4},
			want: func(min *data.Data[int, string], err error) {
				is.Nil(err)
				is.Equal(min.Value, "value1")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range tt.keys {
				mh.Insert(key, fmt.Sprintf("value%d", key))
			}
			min, err := mh.FindMin()
			tt.want(min, err)
		})
	}
}

func Test_DeleteMin(t *testing.T) {
	is := assert.New(t)
	mh := NewMinHeap[int, string]()

	tests := []struct {
		name string
		keys []int
		want func(*data.Data[int, string], error)
	}{
		{
			name: "DeleteMin from an empty heap",
			keys: []int{},
			want: func(min *data.Data[int, string], err error) {
				is.Error(err, fmt.Errorf("empty heap"))
				is.Nil(min)
			},
		},
		{
			name: "DeleteMin from a non-empty heap",
			keys: []int{3, 1, 5, 2, 4, 7},
			want: func(min *data.Data[int, string], err error) {
				is.Nil(err)
				is.Equal(min.Value, "value1")
			},
		},
		{
			name: "DeleteMin from a non-empty heap",
			want: func(min *data.Data[int, string], err error) {
				is.Nil(err)
				is.Equal(min.Value, "value2")
			},
		},
		{
			name: "DeleteMin from a non-empty heap",
			want: func(min *data.Data[int, string], err error) {
				is.Nil(err)
				is.Equal(min.Value, "value3")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range tt.keys {
				mh.Insert(key, fmt.Sprintf("value%d", key))
			}
			min, err := mh.DeleteMin()
			tt.want(min, err)
		})
	}
}
