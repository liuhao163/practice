package main

import "fmt"

type Pkg struct {
	MaxW int
	Pkg  []int
}

func (pkg *Pkg) f(w int, n int) {
	if w == pkg.MaxW || n == len(pkg.Pkg) {
		fmt.Printf("w=%d,maxW=%d,n=%d,maxN=%d\n", w, pkg.MaxW, n, len(pkg.Pkg))
		return
	}

	pkg.f(w, n+1)
	if (w+pkg.Pkg[n] <= pkg.MaxW) {
		pkg.f(w+pkg.Pkg[n], n+1)
	}
}
