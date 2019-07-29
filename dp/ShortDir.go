package dp

var Triangle = [][]int{{5}, {7, 8}, {2, 3, 4}, {4, 9, 6, 1}, {2, 7, 9, 4, 5}}

/**
坐标，只能往做或者往右走，左[i-1][j]-->[i][j],右[i-1][j]-->[i][j+1]
*/
func ShortDir(arr [][]int) int {
	state := make([][]int, len(arr))
	for i := range arr {
		state[i] = make([]int, len(arr))
		for j := range state[i] {
			state[i][j] = -1
		}
	}

	n := len(arr)
	state[0][0] = arr[0][0]

	for i := 1; i < n; i++ {

		for j := 0; j < len(arr[i])-1; j++ {
			if state[i-1][j] > 0 {
				//左走
				state[i][j] = arr[i][j] + state[i-1][j]
				//右走
				state[i][j+1] = arr[i][j+1] + state[i-1][j]
			}
		}
	}

	min := int(^uint(0) >> 1)

	for j := 0; j < len(arr); j++ {
		if state[n-1][j] != -1 && state[n-1][j] < min {
			min = state[n-1][j]
		}
	}

	if min == int(^uint(0)>>1) {
		return -1
	}

	return min
}
