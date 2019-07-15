package main

func BM(s []byte, p []byte) int {
	//bcMap := genBC(p)
	//good suffix
	suffix, prefix := genGS(p)
	for i := 0; i < len(s); {
		j := len(p) - 1
		for ; j >= 0; j-- {
			if s[i+j] != p[j] {
				break
			}

		}

		if j < 0 {
			return i
		}

		//x := j - getBC(s[i+j], bcMap)
		x := -10000

		y := moveGS(j, len(p), suffix, prefix)

		if x < y {
			i = i + y
		} else {
			i = i + x
		}

	}

	return -1
}

func moveGS(j int, pLen int, suffix []int, prefix []bool) int {
	k := pLen - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	} else {
		for r := j + 2; j < pLen-1; j++ {
			if prefix[pLen-1-r+1] {
				return r
			}
		}
	}

	return pLen
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
		for j >= 0 && p[j] >= p[len(p)-1-k] {
			k++
			suffix[k] = j
			j--
		}

		if j < 0 {
			prefix[k] = true
		}
	}

	return suffix, prefix
}
