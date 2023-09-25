package sortedarray

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/sets/data"
	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
)

func TestSortedArray_Find(t *testing.T) {
	is := assert.New(t)

	testData := make([]*data.Data[int, string], 6)
	for i := 0; i < 6; i++ {
		testData[i] = &data.Data[int, string]{
			Key:   i,
			Value: randomdata.Alphanumeric(5),
		}
	}
	testSortedArr := NewSortedArray[int, string](testData...)
	val := testSortedArr.GetLenght() - 1
	fmt.Println(val)
	type args struct {
		key int
	}
	tests := []struct {
		name        string
		sortedarray *SortedArray[int, string]
		args        args
		want        func(string, error)
	}{
		{
			name:        "find random (key 3) element",
			sortedarray: testSortedArr,
			args: args{
				key: 5,
			},
			want: func(got string, err error) {
				is.Nil(err)
				is.Equal(got, testSortedArr.array[5].GetValue())
			},
		},
		{
			name:        "find in empty array",
			sortedarray: NewSortedArray[int, string](),
			args: args{
				key: 0,
			},
			want: func(got string, err error) {
				is.NotNil(err)
				is.Error(err, fmt.Errorf("empty array"))
				is.Empty(got)
			},
		},
		{
			name:        "find first element",
			sortedarray: testSortedArr,
			args: args{
				key: testSortedArr.array[0].Key,
			},
			want: func(got string, err error) {
				is.Nil(err)
				is.Equal(got, testSortedArr.array[0].GetValue())
			},
		},
		{
			name:        "find last element",
			sortedarray: testSortedArr,
			args: args{
				key: testSortedArr.array[testSortedArr.GetLenght()-1].Key,
			},
			want: func(got string, err error) {
				is.Nil(err)
				max, err := testSortedArr.FindMax()
				is.Nil(err)
				is.Equal(got, max)
			},
		},
		{
			name:        "key less than min",
			sortedarray: testSortedArr,
			args: args{
				key: -1,
			},
			want: func(got string, err error) {
				is.NotNil(err)
				is.Error(err, fmt.Errorf("key: %v not found", -1))
			},
		},
		{
			name:        "key greater than max",
			sortedarray: testSortedArr,
			args: args{
				key: len(testSortedArr.array),
			},
			want: func(got string, err error) {
				is.NotNil(err)
				is.Error(err, fmt.Errorf("key: %v not found", testSortedArr.GetLenght()))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.sortedarray.Find(tt.args.key)
			tt.want(data, err)
		})
	}
}

func TestSortedArray_Insert(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name    string
		setData []*data.Data[int, string]
		key     int
		value   string
		want    func(*SortedArray[int, string], error)
	}{
		{
			name:    "insert into an empty array",
			setData: []*data.Data[int, string]{},
			key:     1,
			value:   "one",
			want: func(sa *SortedArray[int, string], err error) {
				is.Nil(err)
				is.Equal(sa.GetLenght(), int32(1))
				is.Equal(sa.GetCapacity(), int32(2))
				is.True(sa.IsSorted())
				v, err := sa.Find(1)
				is.Equal(v, "one")
				is.Nil(err)
			},
		},
		{
			name: "insert into a non-empty array",
			setData: []*data.Data[int, string]{
				data.NewData(1, "one"),
				data.NewData(3, "three"),
			},
			key:   2,
			value: "two",
			want: func(sa *SortedArray[int, string], err error) {
				is.Nil(err)
				is.Equal(sa.GetLenght(), int32(3))
				is.True(sa.IsSorted())
				v, err := sa.Find(2)
				is.Equal(v, "two")
				is.Nil(err)
				min, err := sa.FindMin()
				is.Nil(err)
				is.Equal(min, "one")
				max, err := sa.FindMax()
				is.Nil(err)
				is.Equal(max, "three")
			},
		},
		{
			name: "insert with duplicate key",
			setData: []*data.Data[int, string]{
				data.NewData(1, "one"),
				data.NewData(2, "two"),
				data.NewData(3, "three"),
			},
			key:   2,
			value: "duplicate",
			want: func(sa *SortedArray[int, string], err error) {
				is.Nil(err)
				is.Equal(sa.GetLenght(), int32(3))
				is.True(sa.IsSorted())
				v, err := sa.Find(2)
				is.Equal(v, "duplicate")
				is.Nil(err)
			},
		},
		{
			name: "insert empty key",
			setData: []*data.Data[int, string]{
				data.NewData(1, "one"),
			},
			key:   0,
			value: "zero",
			want: func(sa *SortedArray[int, string], err error) {
				is.Error(err, fmt.Errorf("empty key"))
				is.Equal(sa.GetLenght(), int32(1))
				is.Equal(sa.GetCapacity(), int32(2))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := NewSortedArray[int, string](tt.setData...)
			err := sa.Insert(tt.key, tt.value)
			tt.want(sa, err)
		})
	}
}

func TestSortedArray_Delete(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name    string
		setData []*data.Data[int, string]
		key     int
		want    func(*SortedArray[int, string], error)
	}{
		{
			name: "delete existing key",
			setData: []*data.Data[int, string]{
				data.NewData(1, "one"),
				data.NewData(2, "two"),
				data.NewData(3, "three"),
			},
			key: 2,
			want: func(sa *SortedArray[int, string], err error) {
				is.Nil(err)
				v, err := sa.Find(2)
				is.Empty(v)
				is.Equal(err.Error(), "key: 2 not found")
				is.Equal(sa.GetLenght(), int32(2))
				is.True(sa.IsSorted())
			},
		},
		{
			name: "delete non-existing key",
			setData: []*data.Data[int, string]{
				data.NewData(1, "one"),
				data.NewData(2, "two"),
				data.NewData(3, "three"),
			},
			key: 4,
			want: func(sa *SortedArray[int, string], err error) {
				is.Error(err)
				is.Equal(err.Error(), "key: 4 not found")
				is.Equal(sa.GetLenght(), int32(3))
				is.True(sa.IsSorted())
			},
		},
		{
			name:    "delete from empty array",
			setData: []*data.Data[int, string]{},
			key:     1,
			want: func(sa *SortedArray[int, string], err error) {
				is.Error(err)
				is.Equal(err.Error(), "empty array")
				is.Equal(sa.GetLenght(), int32(0))
				is.True(sa.IsSorted())
			},
		},
		{
			name: "delete last element",
			setData: []*data.Data[int, string]{
				data.NewData(1, "one"),
				data.NewData(3, "three"),
			},
			key: 3,
			want: func(sa *SortedArray[int, string], err error) {
				is.Nil(err)
				v, err := sa.Find(3)
				is.Empty(v)
				is.Equal(err.Error(), "key: 3 not found")
				is.Equal(sa.GetLenght(), int32(1))
				is.True(sa.IsSorted())
			},
		},
		{
			name: "delete first element",
			setData: []*data.Data[int, string]{
				data.NewData(1, "one"),
				data.NewData(2, "two"),
				data.NewData(3, "three"),
				data.NewData(4, "four"),
			},
			key: 1,
			want: func(sa *SortedArray[int, string], err error) {
				is.Nil(err)
				v, err := sa.Find(1)
				is.Empty(v)
				is.Equal(err.Error(), "key: 1 not found")
				is.Equal(sa.GetLenght(), int32(3))
				is.True(sa.IsSorted())
				is.Equal(sa.array[0].GetValue(), "two")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := NewSortedArray[int, string](tt.setData...)
			_, err := sa.Delete(tt.key)
			tt.want(sa, err)
		})
	}
}

func TestSortedArray_FindMin_FindMax(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name    string
		setData []*data.Data[int, string]
		test    func(*SortedArray[int, string])
	}{
		{
			name:    "find min & max in an empty array",
			setData: []*data.Data[int, string]{},
			test: func(sa *SortedArray[int, string]) {
				min, err := sa.FindMin()
				is.Error(err, fmt.Errorf("empty array"))
				is.Equal(min, "")

				max, err := sa.FindMax()
				is.Error(err, fmt.Errorf("empty array"))
				is.Equal(max, "")
			},
		},
		{
			name: "find min & max in a non-empty array",
			setData: []*data.Data[int, string]{
				data.NewData(1, "one"),
				data.NewData(3, "three"),
				data.NewData(2, "two"),
			},
			test: func(sa *SortedArray[int, string]) {
				min, err := sa.FindMin()
				is.Nil(err)
				is.Equal(min, "one")

				max, err := sa.FindMax()
				is.Nil(err)
				is.Equal(max, "three")
			},
		},
		{
			name: "find min & max in an array  single element array",
			setData: []*data.Data[int, string]{
				data.NewData(5, "five"),
			},
			test: func(sa *SortedArray[int, string]) {
				min, err := sa.FindMin()
				is.Nil(err)
				max, err := sa.FindMax()
				is.Nil(err)

				is.Equal(min, max)
			},
		},
		{
			name: "find min & max in an array with duplicate keys",
			setData: []*data.Data[int, string]{
				data.NewData(5, "five"),
				data.NewData(1, "duplicate_one"),
				data.NewData(3, "three"),
				data.NewData(5, "duplicate_five"),
				data.NewData(1, "one"),
			},
			test: func(sa *SortedArray[int, string]) {
				min, err := sa.FindMin()
				is.Nil(err)
				is.Equal(min, "duplicate_one")
				max, err := sa.FindMax()
				is.Nil(err)
				is.Equal(max, "duplicate_five")
			},
		},
		{
			name: "find min & max in an array with negative keys",
			setData: []*data.Data[int, string]{
				data.NewData(-5, "minusfive"),
				data.NewData(-1, "minusone"),
				data.NewData(-10, "minusten"),
			},
			test: func(sa *SortedArray[int, string]) {
				min, err := sa.FindMin()
				is.Nil(err)
				is.Equal(min, "minusten")
				max, err := sa.FindMax()
				is.Nil(err)
				is.Equal(max, "minusone")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := NewSortedArray[int, string](tt.setData...)
			tt.test(sa)
		})
	}
}
