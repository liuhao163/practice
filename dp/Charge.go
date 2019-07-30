package dp

//纸币找零动态规划求解
func ChargeDP(money []int, sum int) int {
	//初始化状态数组 state[i][j]，纸币数，最多是sum，j是当前状态的金额
	state := make([][]bool, sum)
	for i := range state {
		state[i] = make([]bool, sum+1)
		for j := range state[i] {
			state[i][j] = false
		}
	}
	state[0][0] = true

	//min_charge(i,j)=state[i][j+max(money[0] money[1],mongey[2]])==sum-->i
	for i := 1; i < sum; i++ {
		for j := 0; j < sum; j++ {
			//从上一次状态开始推导，当前这次最大的面额然后给纸币数+1
			if state[i-1][j] {
				max := -100
				for k := 0; k < len(money); k++ {
					//小于等于总和，然后取最大值
					if j+money[k] <= sum && max < j+money[k] {
						max = j + money[k]
					}
				}

				//state[i][max]最大值获取最优解
				state[i][max] = state[i-1][j]

				if max == sum {
					return i
				}
			}
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
