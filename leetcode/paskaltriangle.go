package leetcode

import "fmt"

func Print(triangle [][]int) {
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			fmt.Printf("%v ", triangle[i][j])
		}
		fmt.Println()
	}
}

func CreateTriangle(i int) [][]int {
	triangle := make([][]int, 0)
	return doCreatTriangle(triangle, i)
}

func doCreatTriangle(triangle [][]int, i int) [][]int {
	if i == 0 {
		row := make([]int, 1)
		row[0] = 1
		triangle = append(triangle, row)
		return triangle
	}

	triangle = doCreatTriangle(triangle, i-1)

	lastRow := triangle[i-1]
	curRow := make([]int, i+1)
	for col := 0; col <= i; col++ {
		v := 0
		if col == 0 || col == i {
			v = 1
		} else {
			v = lastRow[col] + lastRow[col-1]
		}
		curRow[col] = v
	}
	triangle = append(triangle, curRow)
	return triangle
}
