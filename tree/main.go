package main

import (
	"fmt"
)

func main() {
	trie := NewTrie()
	fmt.Println(*trie)

	s := "liuhao"
	trie.Insert(s)
}
