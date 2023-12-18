package trienode

type Node struct {
	Edges   map[rune]*Node
	WordEnd bool
}

func NewTrieNode() *Node {
	return &Node{
		Edges:   make(map[rune]*Node),
		WordEnd: false,
	}
}

func (n *Node) IsWordEnd() bool {
	return n.WordEnd
}
