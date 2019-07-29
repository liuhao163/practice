package dp

var Pkg = []int{2, 2, 4, 6, 3}
var Value = []int{3, 4, 8, 9, 6}

/**
背包0-1问题升级版，
求解一组物体，pkg代表重量，它的价值是value，装n个物体，在w的限制下，如何使价值最大
*/
func Knapsack(pkg []int, value []int, n int, w int) int {
	//初始化state数组长队[n][w+1]
	state := make([][]int, n)
	for i := 0; i < n; i++ {
		state[i] = make([]int, w+1)
		for j := range state[i] {
			state[i][j] = -1
		}
	}

	//0元素特殊处理
	state[0][0] = 0
	if pkg[0] <= w {
		state[0][pkg[0]] = value[0]
	}

	//从第一个物体开始考察，开始状态转移
	for i := 1; i < n; i++ {
		//不放进背包中，i个物体是i-1个物体的重量j,value是state[i][j]值
		for j := 0; j < w; j++ {
			//如果上一个物体有值
			if state[i-1][j] > -1 {
				state[i][j] = state[i-1][j]
			}
		}

		//放进背包中，i个物体是i-1个物体的重量j+pkg[i],value是state[i-1][j] + value[i]
		//这里求解的是最优解，所以d当v大于当前重量的时候保留v。
		for j := 0; j <= w-pkg[i]; j++ {
			if state[i-1][j] > -1 {
				v := state[i-1][j] + value[i]
				cw := j + pkg[i]
				if v > state[i][cw] {
					state[i][cw] = v
				}
			}
		}
	}

	//遍历获取最大值
	max := -1
	for j := 0; j <= w; j++ {
		if state[n-1][j] > max {
			max = state[n-1][j]
		}
	}
	return max
}
