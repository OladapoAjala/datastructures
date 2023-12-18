package trie

import "github.com/OladapoAjala/datastructures/trie/trienode"

type Trie struct {
	Root *trienode.Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: trienode.NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	curr := t.Root
	for _, r := range word {
		if _, ok := curr.Edges[r]; !ok {
			curr.Edges[r] = trienode.NewTrieNode()
		}
		curr = curr.Edges[r]
	}
	curr.WordEnd = true
}

func (t *Trie) Search(word string) bool {
	if word == "" {
		return false
	}

	curr := t.Root
	for _, r := range word {
		if _, ok := curr.Edges[r]; !ok {
			return false
		}
		curr = curr.Edges[r]
	}
	return curr.IsWordEnd()
}
