package bitmap

type BitMap struct {
	Array []int
	Count int
	Base  int
}

func NewBitMap(count int, base int) BitMap {
	return BitMap{Array: make([]int, count/base+1), Count: count, Base: base}
}

func (this *BitMap) Set(v int) {
	idx := v / this.Base
	bit := uint(v % this.Base)
	this.Array[idx] |= 1 << bit
}

func (this *BitMap) Exists(v int) bool {
	idx := v / this.Base
	bit := uint(v % this.Base)

	return (this.Array[idx] & (1 << bit)) != 0
}
