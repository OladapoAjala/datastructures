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

func Test_InsertFirst(t *testing.T) {
	is := assert.New(t)

	type args struct {
		data string
	}
	tests := []struct {
		name         string
		dynamicarray *DynamicArray[string]
		args         args
		want         func(*DynamicArray[string], error)
	}{
		{
			name:         "insert first data in empty dynamicarray",
			dynamicarray: NewDynamicArray[string](),
			args: args{
				data: "a",
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.True(da.Contains("a"))
				data, err := da.GetData(0)
				is.Nil(err)
				is.Equal(data, "a")
			},
		},
		{
			name:         "replace first data in simple dynamicarray",
			dynamicarray: NewDynamicArray[string]("a", "b", "c"),
			args: args{
				data: "d",
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.False(da.Contains("a"))
				data, err := da.GetData(0)
				is.Nil(err)
				is.Equal(data, "d")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dynamicarray.InsertFirst(tt.args.data)
			tt.want(tt.dynamicarray, err)
		})
	}
}

func Test_InsertLast(t *testing.T) {
	is := assert.New(t)

	type args struct {
		data string
	}
	tests := []struct {
		name         string
		dynamicarray *DynamicArray[string]
		args         args
		want         func(*DynamicArray[string], error)
	}{
		{
			name:         "insert last data in empty dynamicarray",
			dynamicarray: NewDynamicArray[string](),
			args: args{
				data: "a",
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.True(da.Contains("a"))
				data, err := da.GetData(da.Size() - 1)
				is.Nil(err)
				is.Equal(data, "a")
			},
		},
		{
			name:         "replace last data in simple dynamicarray",
			dynamicarray: NewDynamicArray[string]("a", "b", "c"),
			args: args{
				data: "d",
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.False(da.Contains("c"))
				data, err := da.GetData(da.Size() - 1)
				is.Nil(err)
				is.Equal(data, "d")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dynamicarray.InsertLast(tt.args.data)
			tt.want(tt.dynamicarray, err)
		})
	}
}

func Test_Delete(t *testing.T) {
	is := assert.New(t)

	type args struct {
		index int32
	}

	tests := []struct {
		name         string
		dynamicarray *DynamicArray[string]
		args         args
		want         func(*DynamicArray[string], error)
	}{
		{
			name:         "delete from empty dynamicarray",
			dynamicarray: NewDynamicArray[string](),
			args: args{
				index: 0,
			},
			want: func(da *DynamicArray[string], err error) {
				is.Error(err)
				is.Equal(err.Error(), "cannot remove from empty array")
				is.False(da.Contains("a"))
			},
		},
		{
			name:         "delete from simple dynamicarray",
			dynamicarray: NewDynamicArray[string]("a", "b", "c"),
			args: args{
				index: 1,
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.False(da.Contains("b"))
				val, err := da.GetData(0)
				is.Nil(err)
				is.Equal(val, "a")
				val, err = da.GetData(1)
				is.Nil(err)
				is.Equal(val, "c")
				val, err = da.GetData(2)
				is.Equal(err.Error(), "index out of range")
				is.Empty(val)
			},
		},
		{
			name:         "delete from invalid index",
			dynamicarray: NewDynamicArray[string]("a", "b", "c"),
			args: args{
				index: 15,
			},
			want: func(da *DynamicArray[string], err error) {
				is.Error(err)
				is.Equal(err.Error(), "index out of range")
				is.False(da.Contains("f"))
			},
		},
		{
			name:         "delete from dynamicarray with single element",
			dynamicarray: NewDynamicArray[string]("x"),
			args: args{
				index: 0,
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.False(da.Contains("x"))
				is.Equal(da, NewDynamicArray[string]())
			},
		},
		{
			name:         "delete from dynamicarray with capacity reduction",
			dynamicarray: NewDynamicArray[string]("a", "b", "c", "d", "e", "f", "g", "h"),
			args: args{
				index: 0,
			},
			want: func(da *DynamicArray[string], err error) {
				is.Nil(err)
				is.False(da.Contains("a"))

				for i := 1; i <= 3; i++ {
					err = da.Delete(0)
					is.Nil(err)
				}
				is.Equal(da, NewDynamicArray[string]("e", "f", "g", "h"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dynamicarray.Delete(tt.args.index)
			tt.want(tt.dynamicarray, err)
		})
	}
}
