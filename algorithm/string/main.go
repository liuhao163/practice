package main

import "fmt"

func main() {
	s := "ababaeababacdbac"
	p := "ababacd"
	//idx := BM(s, p)
	//fmt.Println(idx)
	//fmt.Printf(" %d,%s\n", idx, s[idx:idx+len(p)])

	idx2 := KMP(s,p)
	fmt.Println(idx2)
 	fmt.Printf("kmp %d,%s\n", idx2, s[idx2:idx2+len(p)])

}
