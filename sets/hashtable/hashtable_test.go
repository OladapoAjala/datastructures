package hashtable

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/sets/data"
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
				is.Equal(ht.Table[1].Head.Data.Key, "key1")
				is.Equal(ht.Table[1].Head.Data.Value, "value1")
				is.Nil(ht.Table[1].Head.Next)
			},
		},
		{
			name: "replace value of previous key",
			args: args{
				key:   "key1",
				value: true,
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[1].Head.Data.Key, "key1")
				is.Equal(ht.Table[1].Head.Data.Value, true)
				is.Nil(ht.Table[1].Head.Next)
			},
		},
		{
			name: "insert multiple key-value pairs",
			args: args{
				key:   "key0",
				value: "value2",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[0].Head.Data.Key, "key0")
				is.Equal(ht.Table[0].Head.Data.Value, "value2")
				is.Nil(ht.Table[0].Head.Next)
			},
		},
		{
			name: "insert key-value pairs with different data types and resize table",
			args: args{
				key:   "key3",
				value: 100,
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[5].Head.Data.Key, "key3")
				is.Equal(ht.Table[5].Head.Data.Value, 100)
				is.Nil(ht.Table[5].Head.Next)

				is.EqualValues(ht.GetCapacity(), 6)
				is.EqualValues(ht.GetSize(), 3)
				is.EqualValues(ht.GetThreshold(), 4)
			},
		},
		{
			name: "insert already existing key-value pairs",
			args: args{
				key:   "key3",
				value: 100,
			},
			want: func(ht *HashTable[string], err error) {
				is.Error(fmt.Errorf("key: key3, value: 100 already in hash table"))
				is.Equal(ht.Table[5].Head.Data.Key, "key3")
				is.Equal(ht.Table[5].Head.Data.Value, 100)
				is.Nil(ht.Table[5].Head.Next)
			},
		},
	}

	hashTable := NewHashTable[string](3)
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

func Test_Delete(t *testing.T) {
	is := assert.New(t)

	type args struct {
		key string
	}
	tests := []struct {
		name  string
		setup func(ht *HashTable[string])
		args  args
		want  func(*HashTable[string], error)
	}{
		{
			name: "delete from empty hashtable",
			args: args{
				key: "key1",
			},
			want: func(ht *HashTable[string], err error) {
				is.Error(fmt.Errorf("key nonexistent not found in hashtable"))
			},
		},
		{
			name: "delete existing key",
			setup: func(ht *HashTable[string]) {
				err := ht.Insert("key1", "value1")
				is.Nil(err)
			},
			args: args{
				key: "key1",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				tmp := data.NewEntry("key1", "value1")
				pos := tmp.GetHash() % uint32(3)
				isPresent, entry := ht.contains(tmp, pos)
				is.Nil(entry)
				is.False(isPresent)
			},
		},
		{
			name: "delete non-existent key",
			setup: func(ht *HashTable[string]) {
				_ = ht.Insert("key2", "value2")
			},
			args: args{
				key: "nonexistent",
			},
			want: func(ht *HashTable[string], err error) {
				is.Error(fmt.Errorf("key nonexistent not found in hashtable"))
			},
		},
		{
			name: "delete key with collision",
			setup: func(ht *HashTable[string]) {
				err := ht.Insert("key1", 100)
				is.Nil(err)
				err = ht.Insert("key2", true)
				is.Nil(err)
			},
			args: args{
				key: "key2",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)

				tmp := data.NewEntry("key2", true)
				pos := tmp.GetHash() % uint32(3)
				isPresent, entry := ht.contains(tmp, pos)
				is.Nil(entry)
				is.False(isPresent)

				tmp = data.NewEntry("key1", 100)
				pos = tmp.GetHash() % uint32(3)
				isPresent, entry = ht.contains(tmp, pos)
				is.Equal(entry.Key, "key1")
				is.Equal(entry.Value, 100)
				is.True(isPresent)
			},
		},
		{
			name: "delete last key in collision list",
			args: args{
				key: "key1",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)

				tmp := data.NewEntry("key1", 100)
				pos := tmp.GetHash() % uint32(3)
				isPresent, entry := ht.contains(tmp, pos)
				is.Nil(entry)
				is.False(isPresent)
			},
		},
	}

	hashTable := NewHashTable[string](3)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(hashTable)
			}

			err := hashTable.Delete(tt.args.key)
			tt.want(hashTable, err)
		})
	}
}
