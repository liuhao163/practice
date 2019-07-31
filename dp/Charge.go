package dp

//纸币找零动态规划求解
func ChargeDP(money []int, sum int) int {
	state := make([][]bool, sum)

	for i := range state {
		state[i] = make([]bool, sum+1)
	}

	state[0][0] = true

	//方程 min[i,j]=state[i][max(j+money[0],j+money[1],j+money[2]+j+...)]
	for i := 1; i < sum; i++ {
		for j := 0; j < sum; j++ {
			if state[i-1][j] {
				for _, m := range money {
					if j+m <= sum {
						state[i][j+m] = true
					}
				}
			}
		}
	}

	for i := 0; i < len(state); i++ {
		if state[i][sum] {
			return i
		}
	}

	return -1
}

var MinVal = int(^uint(0) >> 1)

func ChargeRC(c int, sum int, money []int, min int) {

	var v int
	for _, v = range money {
		if c+v <= sum {
			ChargeRC(c+v, sum, money, min+1)
		}
	}

	if c == sum {
		if min < MinVal {
			MinVal = min
		}
		return
	}
}
