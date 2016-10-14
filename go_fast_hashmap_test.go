package go_fast_hashmap

import (
	"testing"
	"math/rand"
	"time"
)

func BenchmarkFindNextPrime(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		testNum := uint64(r.Int63n(int64(100000000)))
		b.StartTimer()

		nextPrime(testNum)
	}
}
