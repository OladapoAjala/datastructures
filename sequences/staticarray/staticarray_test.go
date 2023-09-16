package staticarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetData(t *testing.T) {
	is := assert.New(t)

	type args struct {
		index int32
	}
	tests := []struct {
		name        string
		staticarray *StaticArray[string]
		args        args
		want        func(string, error)
	}{
		{
			name:        "get data from empty staticarray",
			staticarray: NewStaticArray[string](0),
			args: args{
				index: 0,
			},
			want: func(data string, err error) {
				is.Empty(data)
				is.Error(err, "index out of range")
			},
		},
		{
			name:        "get data at index 2 from staticarray",
			staticarray: NewStaticArray[string](3, "a", "b", "c"),
			args: args{
				index: 2,
			},
			want: func(data string, err error) {
				is.Nil(err)
				is.Equal(data, "c")
			},
		},
		{
			name:        "get data from array with size lesser than input data",
			staticarray: NewStaticArray[string](3, "a", "b", "c", "d", "e"),
			args: args{
				index: 4,
			},
			want: func(data string, err error) {
				is.Empty(data)
				is.Error(err, "index out of range")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.staticarray.GetData(tt.args.index)
			tt.want(data, err)
		})
	}
}

func Test_Contains(t *testing.T) {
	is := assert.New(t)

	type args struct {
		data string
	}
	tests := []struct {
		name        string
		staticarray *StaticArray[string]
		args        args
		want        func(bool)
	}{
		{
			name:        "check if data is in empty staticarray",
			staticarray: NewStaticArray[string](0),
			args: args{
				data: "a",
			},
			want: func(isPresent bool) {
				is.False(isPresent)
			},
		},
		{
			name:        "check if data is in simple staticarray",
			staticarray: NewStaticArray[string](3, "a", "b", "c"),
			args: args{
				data: "b",
			},
			want: func(isPresent bool) {
				is.True(isPresent)
			},
		},
		{
			name:        "check if data is in array with size lesser than input data",
			staticarray: NewStaticArray[string](3, "a", "b", "c", "d", "e"),
			args: args{
				data: "e",
			},
			want: func(isPresent bool) {
				is.False(isPresent)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.staticarray.Contains(tt.args.data)
			tt.want(got)
		})
	}
}

func Test_Insert(t *testing.T) {
	is := assert.New(t)

	type args struct {
		index int32
		data  string
	}
	tests := []struct {
		name        string
		staticarray *StaticArray[string]
		args        args
		want        func(*StaticArray[string], error)
	}{
		{
			name:        "insert data in empty staticarray",
			staticarray: NewStaticArray[string](0),
			args: args{
				index: 0,
				data:  "a",
			},
			want: func(sa *StaticArray[string], err error) {
				is.Error(err, "index out of range")
			},
		},
		{
			name:        "insert data in simple staticarray",
			staticarray: NewStaticArray[string](3, "a", "b", "c"),
			args: args{
				index: 1,
				data:  "d",
			},
			want: func(sa *StaticArray[string], err error) {
				is.Nil(err)
				is.False(sa.Contains("b"))
				data, err := sa.GetData(1)
				is.Nil(err)
				is.Equal(data, "d")
			},
		},
		{
			name:        "insert data at index > array size",
			staticarray: NewStaticArray[string](5, "a", "b", "c", "d", "e"),
			args: args{
				data:  "f",
				index: 6,
			},
			want: func(sa *StaticArray[string], err error) {
				is.Error(fmt.Errorf("index out of range"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.staticarray.Insert(tt.args.index, tt.args.data)
			tt.want(tt.staticarray, err)
		})
	}
}

func Test_Delete(t *testing.T) {
	is := assert.New(t)

	type args struct {
		index int32
	}

	tests := []struct {
		name        string
		staticarray *StaticArray[string]
		args        args
		want        func(*StaticArray[string], error)
	}{
		{
			name:        "delete from empty staticarray",
			staticarray: NewStaticArray[string](0),
			args: args{
				index: 0,
			},
			want: func(sa *StaticArray[string], err error) {
				is.Error(err)
				is.Equal(err.Error(), "index out of range")
			},
		},
		{
			name:        "delete from simple staticarray",
			staticarray: NewStaticArray[string](3, "a", "b", "c"),
			args: args{
				index: 1,
			},
			want: func(sa *StaticArray[string], err error) {
				is.Nil(err)
				is.False(sa.Contains("b"))
				is.Equal(sa, NewStaticArray[string](3, "a", "c", ""))
			},
		},
		{
			name:        "delete from invalid index",
			staticarray: NewStaticArray[string](5, "a", "b", "c", "d", "e"),
			args: args{
				index: 6,
			},
			want: func(sa *StaticArray[string], err error) {
				is.Error(err)
				is.Equal(err.Error(), "index out of range")
				// Ensure no changes in the array
				is.Equal(sa, NewStaticArray[string](5, "a", "b", "c", "d", "e"))
			},
		},
		{
			name:        "delete from staticarray with a single element",
			staticarray: NewStaticArray[string](1, "x"),
			args: args{
				index: 0,
			},
			want: func(sa *StaticArray[string], err error) {
				is.Nil(err)
				is.False(sa.Contains("x"))
				is.Equal(sa, NewStaticArray[string](1, ""))
			},
		},
		{
			name:        "delete last element in a non-empty staticarray",
			staticarray: NewStaticArray[string](3, "a", "b", "c"),
			args: args{
				index: 2,
			},
			want: func(sa *StaticArray[string], err error) {
				is.Nil(err)
				is.False(sa.Contains("c"))
				is.Equal(sa, NewStaticArray[string](3, "a", "b", ""))
			},
		},
		{
			name:        "delete first element in a non-empty staticarray",
			staticarray: NewStaticArray[string](3, "a", "b", "c"),
			args: args{
				index: 0,
			},
			want: func(sa *StaticArray[string], err error) {
				is.Nil(err)
				is.False(sa.Contains("a"))
				is.Equal(sa, NewStaticArray[string](3, "b", "c", ""))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.staticarray.Delete(tt.args.index)
			tt.want(tt.staticarray, err)
		})
	}
}
