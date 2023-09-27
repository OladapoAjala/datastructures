package hashtable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	is := assert.New(t)

	type args struct {
		key   string
		value any
	}

	tests := []struct {
		name string
		args args
		want func(*HashTable[string], error)
	}{
		{
			name: "insert a new key-value pair",
			args: args{
				key:   "key1",
				value: "value1",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[3].Head.Data.Key, "key1")
				is.Equal(ht.Table[3].Head.Data.Value, "value1")
				is.Nil(ht.Table[3].Head.Next)
			},
		},
		{
			name: "insert a duplicate key",
			args: args{
				key:   "key1",
				value: true,
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[3].Head.Data.Key, "key1")
				is.Equal(ht.Table[3].Head.Data.Value, "value1")
				is.Equal(ht.Table[3].Head.Next.Data.Key, "key1")
				is.Equal(ht.Table[3].Head.Next.Data.Value, true)
			},
		},
		{
			name: "insert multiple key-value pairs",
			args: args{
				key:   "key2",
				value: "value2",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[0].Head.Data.Key, "key2")
				is.Equal(ht.Table[0].Head.Data.Value, "value2")
				is.Nil(ht.Table[0].Head.Next)
			},
		},
		{
			name: "insert key-value pairs with different data types",
			args: args{
				key:   "key3",
				value: "100",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[1].Head.Data.Key, "key3")
				is.Equal(ht.Table[1].Head.Data.Value, "100")
				is.Nil(ht.Table[1].Head.Next)
			},
		},
		{
			name: "insert already existing key-value pairs",
			args: args{
				key:   "key3",
				value: "100",
			},
			want: func(ht *HashTable[string], err error) {
				is.Error(fmt.Errorf("key: key3, value: 100 already in hash table"))
				is.Equal(ht.Table[1].Head.Data.Key, "key3")
				is.Equal(ht.Table[1].Head.Data.Value, "100")
				is.Nil(ht.Table[1].Head.Next)
			},
		},
	}

	hashTable := NewHashTable[string](10)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := hashTable.Insert(tt.args.key, tt.args.value)
			tt.want(hashTable, err)
		})
	}
}

func TestHashTable_Find(t *testing.T) {
	is := assert.New(t)

	hashTable := NewHashTable[string](10)
	hashTable.Insert("key1", false)
	hashTable.Insert("key2", "value2")
	hashTable.Insert("key3", "value3")
	hashTable.Insert("key4", 0.0223)

	tests := []struct {
		name      string
		searchKey string
		want      func(any, error)
	}{
		{
			name:      "find existing key, string value",
			searchKey: "key2",
			want: func(v any, err error) {
				is.Nil(err)
				is.Equal(v, "value2")
			},
		},
		{
			name:      "find existing key, boolean value",
			searchKey: "key1",
			want: func(v any, err error) {
				is.Nil(err)
				is.Equal(v, false)
			},
		},
		{
			name:      "find existing key, float value",
			searchKey: "key4",
			want: func(v any, err error) {
				is.Nil(err)
				is.Equal(v, 0.0223)
			},
		},
		{
			name:      "find non-existing key",
			searchKey: "key5",
			want: func(v any, err error) {
				is.Error(fmt.Errorf("key key5 not found in hashtable"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := hashTable.Find(tt.searchKey)
			tt.want(value, err)
		})
	}
}
