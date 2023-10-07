package doublehashing

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
		name  string
		args  args
		setup func(*HashTable[string], string)
		want  func(*HashTable[string], error)
	}{
		{
			name: "insert with invalid key",
			args: args{
				key:   "",
				value: "invalid",
			},
			want: func(ht *HashTable[string], err error) {
				is.Error(fmt.Errorf("invalid key"))
				is.EqualValues(ht.GetSize(), 0)
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 0)
				is.EqualValues(ht.GetLoadFactor(), 0)
			},
		},
		{
			name: "insert a new key-value pair",
			args: args{
				key:   "key1",
				value: "value1",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[3].GetKey(), "key1")
				is.Equal(ht.Table[3].GetValue(), "value1")
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 1)
				is.EqualValues(ht.GetLoadFactor(), float32(0.2))
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
				is.Equal(ht.Table[3].GetKey(), "key1")
				is.Equal(ht.Table[3].GetValue(), true)
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 1)
				is.EqualValues(ht.GetLoadFactor(), float32(0.2))
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
				is.Equal(ht.Table[2].GetKey(), "key0")
				is.Equal(ht.Table[2].GetValue(), "value2")
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 2)
				is.EqualValues(ht.GetLoadFactor(), float32(0.4))
			},
		},
		{
			name: "insert with key collision and probing",
			args: args{
				key:   "key4",
				value: []int{1, 9, 9, 9},
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[1].GetKey(), "key4")
				is.Equal(ht.Table[1].GetValue(), []int{1, 9, 9, 9})
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 3)
				is.EqualValues(ht.GetLoadFactor(), float32(0.6))
			},
		},
		{
			name: "insert to trigger resize",
			args: args{
				key:   "resize1",
				value: "value1",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)

				is.Equal(ht.Table[1].GetKey(), "key0")
				is.Equal(ht.Table[1].GetValue(), "value2")
				is.Equal(ht.Table[2].GetKey(), "key1")
				is.Equal(ht.Table[2].GetValue(), true)
				is.Equal(ht.Table[4].GetKey(), "key4")
				is.Equal(ht.Table[4].GetValue(), []int{1, 9, 9, 9})
				is.Equal(ht.Table[5].GetKey(), "resize1")
				is.Equal(ht.Table[5].GetValue(), "value1")

				is.EqualValues(ht.GetCapacity(), 7)
				is.EqualValues(ht.GetSize(), 4)
				is.EqualValues(ht.GetLoadFactor(), float32(0.5714286))
			},
		},
		{
			name: "insert previously deleted value",
			args: args{
				key:   "key4",
				value: "rebirth",
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Delete(key)
				is.Nil(err)
				is.EqualValues(ht.GetCapacity(), 7)
				is.EqualValues(ht.GetSize(), 3)
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[4].GetKey(), "key4")
				is.Equal(ht.Table[4].GetValue(), "rebirth")
				is.EqualValues(ht.GetCapacity(), 7)
				is.EqualValues(ht.GetSize(), 4)
				is.EqualValues(ht.GetLoadFactor(), float32(0.5714286))
			},
		},
	}

	hashTable := NewHashTable[string](5)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(hashTable, tt.args.key)
			}
			err := hashTable.Insert(tt.args.key, tt.args.value)
			tt.want(hashTable, err)
		})
	}
}

func Test_Find(t *testing.T) {
	is := assert.New(t)

	type args struct {
		key string
	}

	tests := []struct {
		name  string
		args  args
		setup func(*HashTable[string], string)
		want  func(any, error)
	}{
		{
			name: "find with invalid key",
			args: args{
				key: "",
			},
			want: func(value any, err error) {
				is.Error(fmt.Errorf("invalid key"), err)
				is.Nil(value)
			},
		},
		{
			name: "find non-existing key",
			args: args{
				key: "nonexistent",
			},
			want: func(value any, err error) {
				is.Error(fmt.Errorf("key nonexistent not found in hashtable"), err)
				is.Nil(value)
			},
		},
		{
			name: "find existing key -- one content",
			args: args{
				key: "key1",
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Insert(key, "value1")
				is.Nil(err)
			},
			want: func(value any, err error) {
				is.Nil(err)
				is.Equal(value, "value1")
			},
		},
		{
			name: "find existing key -- two contents",
			args: args{
				key: "key0",
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Insert(key, "value2")
				is.Nil(err)
			},
			want: func(value any, err error) {
				is.Nil(err)
				is.Equal(value, "value2")
			},
		},
		{
			name: "find key with collision and probing",
			args: args{
				key: "key4",
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Insert(key, []int{1, 9, 9, 9})
				is.Nil(err)
			},
			want: func(value any, err error) {
				is.Nil(err)
				is.Equal(value, []int{1, 9, 9, 9})
			},
		},
		{
			name: "find key after resize",
			args: args{
				key: "key5", // I need something that clashes with key4 on resize.
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Insert(key, "resizeValue")
				is.Nil(err)
			},
			want: func(value any, err error) {
				is.Nil(err)
				is.Equal(value, "resizeValue")
			},
		},
		{
			name: "find key after deleting different item",
			args: args{
				key: "key5",
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Delete("key4")
				is.Nil(err)
			},
			want: func(value any, err error) {
				is.Nil(err)
				is.Equal(value, "resizeValue")
			},
		},
		{
			name: "find key after deleting same item",
			args: args{
				key: "key5",
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Delete(key)
				is.Nil(err)
			},
			want: func(value any, err error) {
				is.Error(fmt.Errorf("key key5 not found in hashtable"), err)
				is.Nil(value)
			},
		},
	}

	hashTable := NewHashTable[string](5)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(hashTable, tt.args.key)
			}
			result, err := hashTable.Find(tt.args.key)
			tt.want(result, err)
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
		args  args
		setup func(*HashTable[string], string)
		want  func(*HashTable[string], error)
	}{
		{
			name: "delete with invalid key",
			args: args{
				key: "",
			},
			want: func(ht *HashTable[string], err error) {
				is.Error(fmt.Errorf("invalid key"), err)
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 0)
			},
		},
		{
			name: "delete non-existing key",
			args: args{
				key: "nonexistent",
			},
			want: func(ht *HashTable[string], err error) {
				is.Error(fmt.Errorf("key nonexistent not found in hashtable"), err)
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 0)
			},
		},
		{
			name: "delete existing key -- one content",
			args: args{
				key: "key1",
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Insert(key, "value1")
				is.Nil(err)
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 0)
			},
		},
		{
			name: "delete existing key -- two contents",
			args: args{
				key: "key0",
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Insert(key, "value2")
				is.Nil(err)
				err = ht.Insert("key1", "value1")
				is.Nil(err)
				is.EqualValues(ht.GetSize(), 2)

				val, err := ht.Find(key)
				is.Nil(err)
				is.Equal(val, "value2")
				val, err = ht.Find("key1")
				is.Nil(err)
				is.Equal(val, "value1")
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 1)
			},
		},
		{
			name: "delete key with collision and probing",
			args: args{
				key: "key4",
			},
			setup: func(ht *HashTable[string], key string) {
				err := ht.Insert(key, []int{1, 9, 9, 9})
				is.Nil(err)
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.EqualValues(ht.GetCapacity(), 5)
				is.EqualValues(ht.GetSize(), 1)
			},
		},
		{
			name: "delete key after resize",
			args: args{
				key: "key5",
			},
			setup: func(ht *HashTable[string], key string) {
				for i := 0; i < 2; i++ {
					err := ht.Insert(fmt.Sprintf("key_%d", i), fmt.Sprintf("value%d", i))
					is.Nil(err)
				}
				err := ht.Insert(key, "resizeValue")
				is.Nil(err)
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.EqualValues(ht.GetCapacity(), 7)
				is.EqualValues(ht.GetSize(), 3)
			},
		},
		{
			name: "delete key twice",
			args: args{
				key: "key5",
			},
			want: func(ht *HashTable[string], err error) {
				is.Error(fmt.Errorf("key key5 not found in hashtable"), err)
				is.EqualValues(ht.GetCapacity(), 7)
				is.EqualValues(ht.GetSize(), 3)
			},
		},
	}

	hashTable := NewHashTable[string](5)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(hashTable, tt.args.key)
			}
			err := hashTable.Delete(tt.args.key)
			tt.want(hashTable, err)
		})
	}
}
