package main

func BM(s string, p string) int {
	bs := getBS(p)
	suffix, prefix := getSf(p)
	for i := 0; i < len(s)-len(p); {
		j := len(p) - 1
		for ; j >= 0; j-- {
			//j是坏字符
			if s[i+j] != p[j] {
				break
			}
		}

		if j < 0 {
			return i
		}

		//badcode
		x := j - bs[s[i+j]]

		y := 1
		if i > len(p)-1 {
			y = moveSF(j, p, suffix, prefix)
		}

		if x > y {
			i = i + x
		} else {
			i = i + y
		}
	}

	return -1
}
func moveSF(j int, p string, suffix []int, prefix []bool) int {
	l := len(p)
	k := l - 1 - j
	if suffix[k] != -1 {
		return j + 1 - suffix[k]
	}

	for r := j + 2; r < l; r++ {
		if prefix[l-1-r+1] {
			return r
		}
	}

	return l
}

func getSf(p string) ([]int, []bool) {
	len := len(p)
	suffix := make([]int, len)
	prefix := make([]bool, len)
	for i := 0; i < len; i++ {
		j := i
		k := 0
		for j >= 0 && p[i] == p[len-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}

		prefix[k] = j < 0
	}

	return suffix, prefix
}

const asc_size int = 255

func getBS(p string) []int {
	bs := make([]int, asc_size)

	for i := 0; i < asc_size; i++ {
		bs[i] = -1
	}

	for i := 0; i < len(p); i++ {
		bs[p[i]] = i
	}

	return bs;
}
