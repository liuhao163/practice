package main


func KMP(s string, p string) int {
	nexts := getNext(p)

	j := 0
	for i := 0; i < len(s); i++ {
		for ; j > 0 && s[i] != p[j]; {
			j = nexts[j-1]
		}

		if s[i] == p[j] {
			j++
		}

		if j == len(p) {
			//因为s[i] == s[j] 一直到pattern结束所以，下标应该是i-(len(p)-1)
			return i - (len(p) - 1)
		}

	}
	return -1
}

func getNext(p string) []int {
	ret := make([]int, len(p))
	ret[0]=0
	maxLen := 0
	for i := 1; i < len(p); i++ {
		for ; maxLen > 0 && p[maxLen] != p[i]; {
			maxLen = ret[maxLen-1]
		}

		if p[maxLen] == p[i] {
			maxLen++
		}

		ret[i] = maxLen
	}

	return ret;
}
