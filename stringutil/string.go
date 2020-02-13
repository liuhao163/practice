package stringutil

func ReverseStr(str string) string {

	ret := doResverse([]byte(str), 0, len(str)/2-1)
	return string(ret)
}

func doResverse(str []byte, idx int, mid int) []byte {
	if len(str) == 0 || idx >= len(str) {
		return str
	}

	str = doResverse(str, idx+1, mid)
	if idx > mid {
		tmp := str[idx]
		str[idx] = str[len(str)-1-idx]
		str[len(str)-1-idx] = tmp
	}
	return str

}
