package bt

import "fmt"

type Queen struct {
	ret   []int
	total int
}

func NewQueen(total int) Queen {
	ret := make([]int, total)
	for i := range ret {
		ret[i] = -1
	}

	return Queen{ret, total}
}

func (this *Queen) Queen(row int) {
	if row == this.total {
		this.print()
		fmt.Println("!!!!!!!!!!!!!!!!!!!")
		return
	}

	for col := 0; col < this.total; col++ {
		if this.ok(row, col) {
			this.ret[row] = col
			this.Queen(row + 1)
		}
	}
}

func (this *Queen) ok(row int, col int) bool {

	leftUp := col - 1
	rightUp := col + 1
	for i := row - 1; i >= 0; i--{
		if this.ret[i] == col {
			return false
		}
		if leftUp >= 0 && this.ret[i] == leftUp {
			return false
		}
		if rightUp <= this.total && this.ret[i] == rightUp {
			return false
		}
		rightUp = rightUp + 1
		leftUp = leftUp - 1

	}

	return true
}

func (this Queen) print() {
	for i := 0; i < this.total; i++ {
		for j := 0; j < this.total; j++ {
			if this.ret[i] == j {
				fmt.Print("* ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}

}
