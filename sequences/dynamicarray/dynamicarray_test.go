package dynamicarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetData(t *testing.T) {
	is := assert.New(t)

	type args struct {
		index int32
	}
	tests := []struct {
		name         string
		dynamicarray *DynamicArray[string]
		args         args
		want         func(string, error)
	}{
		{
			name:         "get data from empty dynamicarray",
			dynamicarray: NewDynamicArray[string](),
			args: args{
				index: 0,
			},
			want: func(data string, err error) {
				is.Empty(data)
				is.Error(err, "index out of range")
			},
		},
		{
			name:         "get data at index 2 from dynamicarray",
			dynamicarray: NewDynamicArray[string]("a", "b", "c"),
			args: args{
				index: 2,
			},
			want: func(data string, err error) {
				is.Nil(err)
				is.Equal(data, "c")
			},
		},
		{
			name:         "get data from range larger than lenght but lesser than capacity",
			dynamicarray: NewDynamicArray[string]("a", "b", "c", "d", "e"),
			args: args{
				index: 7,
			},
			want: func(data string, err error) {
				is.Empty(data)
				is.Error(err, "index out of range")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.dynamicarray.GetData(tt.args.index)
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
		name         string
		dynamicarray *DynamicArray[string]
		args         args
		want         func(bool)
	}{
		{
			name:         "check if data is in empty dynamicarray",
			dynamicarray: NewDynamicArray[string](),
			args: args{
				data: "a",
			},
			want: func(isPresent bool) {
				is.False(isPresent)
			},
		},
		{
			name:         "check if data is in simple dynamicarray",
			dynamicarray: NewDynamicArray[string]("a", "b", "c"),
			args: args{
				data: "b",
			},
			want: func(isPresent bool) {
				is.True(isPresent)
			},
		},
		{
			name:         "check if null data is in dynamicarray",
			dynamicarray: NewDynamicArray[string]("a", "b", "c"),
			args: args{
				data: "",
			},
			want: func(isPresent bool) {
				is.False(isPresent)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isPresent := tt.dynamicarray.Contains(tt.args.data)
			tt.want(isPresent)
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
		name         string
		dynamicarray *DynamicArray[string]
		args         args
		want         func(*DynamicArray[string], error)
	}{
		{
			name:         "insert data in empty dynamicarray",
			dynamicarray: NewDynamicArray[string](),
			args: args{
				index: 0,
				data:  "a",
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.True(da.Contains("a"))
			},
		},
		{
			name:         "insert data in simple dynamicarray",
			dynamicarray: NewDynamicArray[string]("a", "b", "c"),
			args: args{
				index: 1,
				data:  "d",
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.False(da.Contains("b"))
				is.Equal(da, NewDynamicArray[string]("a", "d", "c"))
			},
		},
		{
			name:         "insert data at index > array capacity",
			dynamicarray: NewDynamicArray[string]("a", "b", "c"),
			args: args{
				data:  "f",
				index: 15,
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.True(da.Contains("f"))
				data, err := da.GetData(15)
				is.Nil(err)
				is.Equal(data, "f")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dynamicarray.Insert(tt.args.index, tt.args.data)
			tt.want(tt.dynamicarray, err)
		})
	}
}
