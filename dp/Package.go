package dp

var Pkg = []int{2, 2, 4, 6, 3}
var Value = []int{2, 2, 4, 6, 3}

func knapsack(pkg []int, value []int, n int, w int) int {
	state := [5][10]int{}

	for _, arrayi := range state {
		for j, _ := range arrayi {
			arrayi[j] = -1
		}
	}

	if pkg[0] <= w {
		state[0][pkg[0]] = value[0]
	}

	for i, arrayi := range state[1:] {
		for j := 0; j < w; j++ {
			if state[i-1][j] > -1 {
				state[i][j] = state[i][j]
			}
		}

		for j := 0; j < w-arrayi[i]; j++ {
			if state[i-1][j] > -1 {
				v := state[i-1][j] + pkg[i]
				cw := j + pkg[i]
				if v > state[i][cw] {
					state[i][cw] = v
				}
			}
		}
	}

	max := -1
	for j := 0; j <= w; j++ {
		if max > state[4][j] {
			max = state[4][j]
		}
	}
	return max
}
