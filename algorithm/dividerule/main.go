package main

import "fmt"

func main() {
	rc := &ReversedCount{A: []int{2, 4, 3, 1, 5, 6}, Num: 0}
	rc.Count(0, len(rc.A)-1)
	fmt.Println(rc.Num)
	fmt.Println(rc.A)
}
