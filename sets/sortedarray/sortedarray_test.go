package sortedarray

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/sets/data"
	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
)

func Test_Find(t *testing.T) {
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
				is.Equal(got, testSortedArr.FindMax())
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
