package main

import "testing"

func lenChain(n int) int {
	c := 1
	for ; n > 1; c++ {
		if n&1 == 0 {
			n >>= 1
		} else {
			n += n<<1 + 1
		}
	}
	return c
}

func euler14() (maxn, maxc int) {
	maxs := 1000000 - 1
	for n := 1; n <= maxs; n++ {
		c := lenChain(n)
		if maxc < c {
			maxn = n
			maxc = c
		}
	}
	return maxn, maxc
}

func BenchmarkEuler14(b *testing.B) {
	for N := 0; N < b.N; N++ {
		euler14()
	}
}
