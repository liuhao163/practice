package main

func KMP(s string, p string) int {
	next := getNext(p);
	j := 0
	for i := 0; i < len(s); i++ {
		for ; j > 0 && s[i] != p[j]; {
			j = next[j-1] + 1
		}

		if s[i] == p[j] {
			j++
		}

		if j == len(p) {
			return i - (len(p) - 1)
		}

	}

	return -1
}

func getNext(p string) []int {
	ret := make([]int, len(p))
	ret[0] = -1
	k := -1
	for i := 1; i < len(p); i++ {
		for ; k != -1 && p[k+1] != p[i]; {
			k = ret[k]
		}

		if p[i] == p[k+1] {
			k++
		}

		ret[i] = k
	}

	return ret
}
