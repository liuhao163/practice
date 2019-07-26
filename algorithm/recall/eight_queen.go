package main

import "fmt"

var ret =[]int{-1,-1,-1,-1,-1,-1,-1,-1}

func CalQueen(row int) {
	if row ==8 {
		printQueen()
		return
	}

	for column := 0; column < 8; column++ {
		if ok(row, column) {
			ret[row] = column
			CalQueen(row + 1)
		}
	}

}

func ok(row int, cloumn int) bool {
	leftup := cloumn - 1
	rightup := cloumn + 1
	for i := row - 1; i >= 0; i-- {
		if ret[i] == cloumn {
			return false
		}

		if leftup>=0 && ret[i] == leftup {
			return false
		}

		if rightup<=8 && ret[i] == rightup {
			return false
		}
		leftup--
		rightup++
	}

	return true
}

func printQueen() {
	for row := 0; row < 8; row++ {
		for clo := 0; clo < 8; clo++ {
			if ret[row]== clo{
				fmt.Print("* ")
			}else{
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
	fmt.Println("-----------------------")
}
