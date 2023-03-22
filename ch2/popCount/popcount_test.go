package popcount

import "testing"

const (
	LoopTime = 10000000
)

func BenchmarkPopCountV1(b *testing.B) {
	b.ResetTimer()
	for i := uint64(0); i < LoopTime; i++ {
		_ = PopCount(i)
	}
}

func BenchmarkPopCountV2(b *testing.B) {
	b.ResetTimer()
	for i := uint64(0); i < LoopTime; i++ {
		_ = PopCountV2(i)
	}
}

func BenchmarkPopCountV3(b *testing.B) {
	b.ResetTimer()
	for i := uint64(0); i < LoopTime; i++ {
		_ = PopCountV3(i)
	}
}

func BenchmarkPopCountV4(b *testing.B) {
	b.ResetTimer()
	for i := uint64(0); i < LoopTime; i++ {
		_ = PopCountV4(i)
	}
}


// goos: linux
// goarch: amd64
// pkg: github.com/chengshusss/goplpractice/ch2/popCount
// cpu: 12th Gen Intel(R) Core(TM) i5-12600KF
// BenchmarkPopCountV1
// BenchmarkPopCountV1-16    	1000000000	         0.001033 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPopCountV2
// BenchmarkPopCountV2-16    	1000000000	         0.03650 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPopCountV3
// BenchmarkPopCountV3-16    	1000000000	         0.06393 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPopCountV4
// BenchmarkPopCountV4-16    	1000000000	         0.05041 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/chengshusss/goplpractice/ch2/popCount	1.265s