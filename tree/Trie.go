package main

import "fmt"

type TrieNode struct {
	Value     string
	Children  []TrieNode
	IsEndWord bool
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
	tmp := trie
	for _, b := range bytes {
		i := b - []byte("a")[0]
		if v := tmp.Children[i].Value; v == "" {
			tmp.Children[i] = TrieNode{Value: string(b), Children: make([]TrieNode, 26)}
		}
		tmp = &tmp.Children[i]
	}
	tmp.IsEndWord = true
}

func (trie *TrieNode) Find(s string) bool {
	bytes := []byte(s)
	tmp := *trie
	for _, b := range bytes {
		i := b - []byte("a")[0]
		if v := tmp.Children[i].Value; v == "" {
			return false
		}
		tmp = tmp.Children[i]
	}
	if tmp.IsEndWord {
		return true
	} else {
		fmt.Printf("this is not word.:%s\n", s)
		return false
	}
}
