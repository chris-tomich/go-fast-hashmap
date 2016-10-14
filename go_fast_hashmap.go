package go_fast_hashmap

import (
	"math"
	"github.com/OneOfOne/xxhash"
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
	Key string
	Value string
	Next *bucket
}

type Hashmap struct {
	buckets []*bucket
	bSize   uint64

	hasher  *xxhash.XXHash64
}

func New(size uint64) *Hashmap {
	bSize := findHashmapPrimeSize(size)

	m := &Hashmap{
		buckets: make([]*bucket, bSize),
		bSize: bSize,
		hasher: xxhash.New64(),
	}

	return m
}

func findMatchingKeyOrLastBucket(key string, b *bucket) (*bucket, bool) {
	if b == nil {
		return nil, false
	} else if b.Next == nil {
		return b, false
	} else if b.Key == key {
		return b, true
	} else {
		return findMatchingKeyOrLastBucket(key, b.Next)
	}
}

func (m *Hashmap) Get(key string) (string, bool) {
	m.hasher.Reset()
	m.hasher.WriteString(key)
	h := m.hasher.Sum64()

	index := h % m.bSize

	b := m.buckets[index]

	last, isMatching := findMatchingKeyOrLastBucket(key, b)

	if isMatching {
		return last.Value, true
	} else {
		return "", false
	}
}

func (m *Hashmap) Set(key string, value string) {
	m.hasher.Reset()
	m.hasher.WriteString(key)
	h := m.hasher.Sum64()

	index := h % m.bSize

	b := m.buckets[index]

	if b == nil {
		m.buckets[index] = &bucket{
			Key: key,
			Value: value,
		}
	} else {
		last, isMatching := findMatchingKeyOrLastBucket(key, b)

		if isMatching {
			last.Value = value
		} else {
			last.Next = &bucket{
				Key: key,
				Value: value,
			}
		}
	}
}
