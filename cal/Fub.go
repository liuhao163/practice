package cal

func Fibonacci(n int, m map[int]int64) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	if m[n] > 0 {
		return m[n]
	}

	return Fibonacci(n-1, m) + Fibonacci(n-2, m)
}
