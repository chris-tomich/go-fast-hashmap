package go_fast_hashmap

import (
	"math"
)

// This appears to be a pretty popular load factor figure (taken from a mixture of the Go built-in map, Mono, and .NET Core)
const loadFactor float64 = 0.7

var primeBasedSizes = []uint64{
									11, 101, 211, 503, 1009, 1511, 2003, 3511, 5003, 7507, 10007, 15013, 20011,
									25013, 50021, 75011, 100003, 125003, 150001, 175003, 200003, 350003, 500009,
									750019, 1000003, 1250003, 1750009, 1500007, 2000003, 3500017, 5000011, 7500013,
									10000019,
								}

func isPrime(n uint64) bool {
	for i := uint64(3); i <= (n+3) / 3; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func nextPrime(start uint64) uint64 {
	num := start

	if num % 2 == 0 {
		num++
	}

	for ; !isPrime(num); num += 2 {
		if num > math.MaxUint64 {
			panic("Requested size is too large.")
		}
	}

	return num
}

func findHashmapPrimeSize(size uint64) uint64 {
	maxSize := uint64(float64(size) / loadFactor)

	for _, prime := range primeBasedSizes {
		if prime > maxSize {
			return prime
		}
	}

	// The dataset is clearly huge so we'll just calculate the next prime beyond the maxSize we were given.
	// This obviously could take a while but clearly the user is expecting this as their size is beyond 10,000,000.
	// The following prime finding algorithm is hugely inadequate but it'll have to do for now.

	largePrime := nextPrime(maxSize)

	return largePrime
}

type bucket struct {
	Key []byte
	Value interface{}
	Next *bucket
}

type Hashmap struct {
	buckets []*bucket
}

func New(size uint64) *Hashmap {
	bSize := findHashmapPrimeSize(size)

	m := &Hashmap{
		buckets: make([]*bucket, bSize),
	}

	return m
}

//func (m *Hashmap) Get(key []byte) (interface{}, bool) {
//
//}
//
//func (m *Hashmap) Set(key []byte, value interface{}) {
//
//}
