package main

import "fmt"

func main() {
	s := "abbbbbbbbbbbbbb"
	p := "bbbb"
	fmt.Print(len(p))
	idx := BM([]byte(s), []byte(p))
	fmt.Printf(" %d,%s\n", idx, s[idx:idx+len(p)])

}
