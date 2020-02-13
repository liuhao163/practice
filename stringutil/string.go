package stringutil

func ReverseStr(str []byte) {
	doReverseStr(str, 0, len(str)/2-1)
}

func doReverseStr(str []byte, idx int, mid int) {
	if len(str) == 0 || idx >= len(str) {
		return
	}
	doReverseStr(str, idx+1, mid)
	if idx > mid {
		tmp := str[idx]
		str[idx] = str[len(str)-1-idx]
		str[len(str)-1-idx] = tmp
	}
}
