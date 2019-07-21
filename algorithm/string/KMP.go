package main

import "fmt"

func KMP(s string, p string) int {
	next := getNexts(p)
	j := 0
	i := 0

	for i < len(s) && j < len(p) {
		if j ==-1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(p) {
		return i-j
	}else{
		return -1
	}

	//
	//for i := 0; i < len(s); i++ {
	//
	//
	//	for ; j > 0 && s[i] != p[j]; {
	//		j = next[j-1] + 1
	//	}
	//
	//	if s[i] == p[j] {
	//		j++
	//	}
	//
	//	if j == len(p) {
	//		return i - len(p) + 1
	//	}
	//}

	//return -1
}

func getNexts(p string) []int {
	nexts := make([]int, len(p))
	nexts[0] = -1
	k := -1
	for j := 0; j < len(p)-1; {
		if k == -1 || p[j] == p[k] {
			j++
			k++
			nexts[j] = k
		} else {
			k = nexts[k]
		}
	}

	fmt.Println(nexts)

	//nexts[0] = -1
	//k := -1
	//for i := 1; i < len(p); i++ {
	//	for ; k >= 0 && p[k+1] != p[i]; {
	//		//ababac  假设i=5(c),这时候是k=2,p[k+1]!=p[5] ，b!=c，
	//		// 这时候找次长可匹配后缀，即次长可匹配前缀，一定在next数组中，这时候next数组最长的是
	//		k = nexts[k]
	//	}
	//
	//	if p[k+1] == p[i] {
	//		k++
	//	}
	//	nexts[i] = k
	//}
	return nexts
}
