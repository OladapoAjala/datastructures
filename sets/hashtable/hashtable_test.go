package hashtable

import (
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
				value: "value2",
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[3].Head.Data.Key, "key1")
				is.Equal(ht.Table[3].Head.Data.Value, "value1")
				is.Equal(ht.Table[3].Head.Next.Data.Key, "key1")
				is.Equal(ht.Table[3].Head.Next.Data.Value, "value2")
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
				value: 100,
			},
			want: func(ht *HashTable[string], err error) {
				is.Nil(err)
				is.Equal(ht.Table[1].Head.Data.Key, "key3")
				is.Equal(ht.Table[1].Head.Data.Value, 100)
				is.Nil(ht.Table[1].Head.Next)
			},
		},
	}

	hashTable := NewHashTable[string]()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := hashTable.Insert(tt.args.key, tt.args.value)
			tt.want(hashTable, err)
		})
	}
}
