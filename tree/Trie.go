package main

type TrieNode struct {
	Value    string
	Children []TrieNode
}

func NewTrie() *TrieNode {
	ret := &TrieNode{Value: "/", Children: make([]TrieNode, 26)}
	return ret
}

func Find(s string) bool {

	return false
}

func (trie *TrieNode) Insert(s string) {
	bytes := []byte(s)
	for _, b := range bytes {
		i := b - []byte("a")[0]
		if v := trie.Children[i].Value; v == "" {
			trie.Children[i] = TrieNode{Value: string(b), Children: make([]TrieNode, 26)}
		}
	}
}
