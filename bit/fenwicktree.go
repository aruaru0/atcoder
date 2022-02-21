package bit

// ** copy from here to your source code **
//
// fenwick tree
//  1. go lang have no generics, so implement int(long long)
//  data type only.
//  2. do not care overflow
//

// BIT :
type BIT struct {
	n    int
	data []int
}

// FenwickTree :
func FenwickTree(n int) *BIT {
	var ret BIT
	ret.n = n
	ret.data = make([]int, n)
	return &ret
}

// Add :
func (b *BIT) Add(p, x int) {
	p++
	for p <= b.n {
		b.data[p-1] += x
		p += p & -p
	}
}

// Sum :
func (b *BIT) Sum(l, r int) int {
	return b.sum(r) - b.sum(l)
}

func (b *BIT) sum(r int) int {
	s := 0
	for r > 0 {
		s += b.data[r-1]
		r -= r & -r
	}
	return s
}
