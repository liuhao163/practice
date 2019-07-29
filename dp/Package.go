package dp

var Pkg = []int{2, 2, 4, 6, 3}
var Value = []int{3, 4, 8, 9, 6}

/**
最小背包问题
*/
func Knapsack(pkg []int, value []int, n int, w int) int {
	state := make([][]int, n)

	for i := 0; i < n; i++ {
		state[i] = make([]int, w+1)
		for j := range state[i] {
			state[i][j] = -1
		}
	}

	state[0][0] = 0
	if pkg[0] <= w {
		state[0][pkg[0]] = value[0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < w; j++ {
			if state[i-1][j] > -1 {
				state[i][j] = state[i-1][j]
			}
		}

		for j := 0; j < w-pkg[i]; j++ {
			if state[i-1][j] > -1 {
				v := state[i-1][j] + value[i]
				cw := j + pkg[i]
				if v > state[i][cw] {
					state[i][cw] = v
				}
			}
		}
	}

	max := -1
	for j := 0; j <= w; j++ {
		if state[n-1][j] > max {
			max = state[n-1][j]
		}
	}
	return max
}
