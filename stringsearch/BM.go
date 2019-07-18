package stringsearch

func BM(s string, p string) int {
	suffix, prefix := initGS(p)
	bc := initBC(p)
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

		x := j - bc[s[i+j]]

		y := 1
		if j < len(p)-1 {
			y = moveGS(j, len(p), suffix, prefix)
		}

		ret := x
		if y > x {
			ret = y
		}
		i = i + ret
	}
	return -1
}

func initBC(pattern string) []int {
	bc := make([]int, 256)
	for r := 0; r < 255; r++ {
		bc[r] = -1
	}

	for idx, v := range []byte(pattern) {
		bc[v] = idx
	}

	return bc
}

func initGS(pattern string) ([]int, []bool) {
	suffix := make([]int, len(pattern))
	prefix := make([]bool, len(pattern))
	for r := 0; r < len(pattern); r++ {
		suffix[r] = -1
	}

	for i := 0; i < len(pattern)-1; i++ {
		j := i
		k := 0
		for j >= 0 && pattern[j] == pattern[len(pattern)-1-k] {
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

func moveGS(j int, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	} else {
		for r := j + 2; r <= m-1; r++ {
			if prefix[m-r] {
				return r
			}
		}
	}

	return m
}
