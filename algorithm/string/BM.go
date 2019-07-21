package main

import "fmt"

func BM(s []byte, p []byte) int {
	bcMap := genBC(p)
	//good suffix
	suffix, prefix := genGS(p)
	for i := 0; i < len(s)-len(p); {
		j := len(p) - 1
		for ; j >= 0; j-- {
			if s[i+j] != p[j] {
				break
			}

		}

		if j < 0 {
			return i
		}

		x := j - getBC(s[i+j], bcMap)
		y := 0
		if j < len(p)-1 {
			y = moveGS(j, len(p), suffix, prefix)
		}

		if x < y {
			i = i + y
		} else {
			i = i + x
		}

	}

	return -1
}

func moveGS(j int, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	for r := j + 2; r <= m-1; r++ {
		if prefix[m-r] {
			return r
		}
	}

	return m
}

const asc_size int = 255

func getBC(bc byte, bcMap map[byte]int) int {
	if ret, ok := bcMap[bc]; ok {
		return ret
	}

	return -1
}

func genBC(p []byte) map[byte]int {
	ret := make(map[byte]int, asc_size)
	for idx, v := range p {
		ret[v] = idx
	}
	return ret
}

func genGS(p []byte) ([]int, []bool) {
	suffix := make([]int, len(p))
	prefix := make([]bool, len(p))

	for idx, _ := range suffix {
		suffix[idx] = -1
	}

	for i := 0; i < len(p)-1; i++ {
		j := i
		k := 0
		for j >= 0 && p[j] == p[len(p)-1-k] {
			k++
			j--
			suffix[k] = j + 1
		}

		if j == -1 {
			prefix[k] = true
		}
	}

	fmt.Println(suffix)

	return suffix, prefix
}
