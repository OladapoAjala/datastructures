package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Trie(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name   string
		words  []string
		search string
		want   func(bool)
	}{
		{
			name:   "Insert and search single word",
			words:  []string{"test"},
			search: "test",
			want: func(b bool) {
				is.True(b)
			},
		},
		{
			name:   "Search non-existent word",
			words:  []string{"hello"},
			search: "world",
			want: func(b bool) {
				is.False(b)
			},
		},
		{
			name:   "Insert multiple words and search",
			words:  []string{"cat", "dog", "bird"},
			search: "dog",
			want: func(b bool) {
				is.True(b)
			},
		},
		{
			name:   "Search empty string",
			words:  []string{"empty"},
			search: "",
			want: func(b bool) {
				is.False(b)
			},
		},
		{
			name:   "Insert and search word with special characters",
			words:  []string{"hello-world"},
			search: "hello-world",
			want: func(b bool) {
				is.True(b)
			},
		},
		{
			name:   "Insert and search word with numbers",
			words:  []string{"abc123"},
			search: "abc123",
			want: func(b bool) {
				is.True(b)
			},
		},
		{
			name:   "Search a prefix of an existing word",
			words:  []string{"prefix"},
			search: "pre",
			want: func(b bool) {
				is.False(b)
			},
		},
		{
			name:   "Insert duplicate words",
			words:  []string{"duplicate", "duplicate"},
			search: "duplicate",
			want: func(b bool) {
				is.True(b)
			},
		},
		{
			name:   "Insert and search case sensitive words",
			words:  []string{"Case", "case"},
			search: "Case",
			want: func(b bool) {
				is.True(b)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewTrie()
			for _, word := range tt.words {
				trie.Insert(word)
			}
			got := trie.Search(tt.search)
			tt.want(got)
		})
	}
}
