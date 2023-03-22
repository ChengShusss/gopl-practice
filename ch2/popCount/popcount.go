package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountV2(x uint64) int {
	r := byte(0)
	for i := 0; i < 8; i++ {
		r += pc[byte(x>>(i*8))]
	}
	return int(r)
}

func PopCountV3(x uint64) int {
	r := 0
	for x > 0 {
		if x%2 != 0 {
			r += 1
		}
		x = x >> 1
	}
	return r
}

func PopCountV4(x uint64) int {
	r := 0
	for x > 0 {
		x = x & (x - 1)
	}
	return r
}
