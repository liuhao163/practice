package main

import (
	"fmt"
)

func main() {
	trie := NewTrie()
	//fmt.Println(*trie)

	trie.Insert("hello")
	trie.Insert("he")
	trie.Insert("her")

	fmt.Println(*trie)

	fmt.Println(trie.Find("her"))
}
