package leetcode

func ReverseString(s string) string {
	str := []byte(s)
	doReverseString(str, 0)
	return string(str)
}

func doReverseString(s []byte, i int) {
	if i == len(s)/2 {
		return
	}

	tmp := s[i]
	s[i] = s[len(s)-1-i]
	s[len(s)-1-i] = tmp

	doReverseString(s, i+1)
}
