package main

type ReversedCount struct {
	Num int
	A   []int
}

func (this *ReversedCount) Count(p int, r int) {
	if p >= r {
		return
	}

	m := (r + p) / 2

	this.Count(p, m)
	this.Count(m+1, r)

	this.merge(p, m, r)
}

func (this *ReversedCount) merge(p int, m int, r int) {
	j := p
	k := m + 1
	tmp := make([]int, r-p+1)
	i := 0
	for j <= m && k <= r {
		if this.A[j] <= this.A[k] {
			tmp[i] = this.A[j]
			j++
			i++
		} else {
			tmp[i] = this.A[k]
			/**
			1.数组分为俩部分分别为前半部分a[p,m]，以及后版部分a[m+1,r]
			2.这时候因为A[K]<A[J]，又因为前半部分的数组是有序的，所以前半部分的数组的剩余部分都是逆序的，统计这部分元素个数即可
			3.计算方法m-j+1,解释：
				j是之前已经有了j个比A[k]小的元素，由于前半数组是有序的，所以从m-j+1个之后都是比a[j]大的元素了。
				+1是因为结果药品包含当前元素（这个地方当时写错了）
			*/
			this.Num += m - j + 1
			k++
			i++
		}
	}

	for j <= m {
		tmp[i] = this.A[j]
		i++
		j++
	}

	for k <= r {
		tmp[i] = this.A[k]
		i++
		k++
	}

	for idx, i := range tmp {
		this.A[p+idx] = i
	}

}
