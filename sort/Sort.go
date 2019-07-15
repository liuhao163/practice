package sort

func SelectSort(a []int) *[]int {
	for i := 0; i < len(a); i++ {
		minIdx := i
		min := a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				minIdx = j
				min = a[j]
			}
		}
		if i != minIdx {
			tmp := a[i]
			a[i] = a[minIdx]
			a[minIdx] = tmp
		}
	}

	return &a
}

func InsertSort(a []int) *[]int {
	for i := 1; i < len(a); i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}

		a[j+1] = v
	}
	return &a
}

func BubbleSort(a []int) *[]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}

	return &a
}

func MergeSort(a []int, s, e int) {
	if s >= e {
		return
	}

	mid := (s + e) / 2
	MergeSort(a, s, mid)
	MergeSort(a, mid+1, e)

	merge(a, s, mid, e)
}
func merge(a []int, s int, mid int, e int) {
	tmp := make([]int, e-s+1)
	k := 0

	i := s
	j := mid + 1

	for i <= mid && j <= e {
		if a[i] < a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}

	for i <= mid {
		tmp[k] = a[i]
		k++
		i++
	}

	for j <= e {
		tmp[k] = a[j]
		k++
		j++
	}

	for x := 0; x < len(tmp); x++ {
		a[s+x] = tmp[x]
	}
}

func QuickSort(a []int, s, e int) {
	if s >= e {
		return
	}

	point := getPoivt(a, s, e)

	QuickSort(a, s, point-1)
	QuickSort(a, point, e)
}
func getPoivt(a []int, s int, e int) int {
	flag := a[e]
	swapIdx := s
	for i := s; i <= e-1; i++ {
		if a[i] <= flag {
			tmp := a[i]
			a[i] = a[swapIdx]
			a[swapIdx] = tmp
			swapIdx++
		}
	}

	tmp := a[swapIdx]
	a[swapIdx] = flag
	a[e] = tmp

	return swapIdx
}
