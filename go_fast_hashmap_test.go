package go_fast_hashmap

import (
	"testing"
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

func TestGoFastHashmap(t *testing.T) {
	m := New(10)

	m.Set("A", "a")
	m.Set("B", "b")
	m.Set("C", "c")
	m.Set("D", "d")

	fmt.Println(m.Get("A"))
	fmt.Println(m.Get("B"))
	fmt.Println(m.Get("C"))
	fmt.Println(m.Get("D"))
}

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

func MakeWord(maxSize int) string {
	var letters = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	wordSize := r.Intn(maxSize);

	if wordSize < 3 {
		wordSize = 3
	}

	isAlpha := r.Intn(10) >= 3

	var word string

	for i := 0; i < wordSize; i++ {
		if isAlpha {
			word = word + letters[r.Intn(26)]
		} else {
			word = word + strconv.Itoa(r.Intn(10))
		}
	}

	return word
}

func GetTwoMatchingSizedSets(size int) ([]string, []string) {
	bagOfWords := make([]string, 0, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		bagOfWords = append(bagOfWords, MakeWord(10))
	}

	wordSubSetLen := len(bagOfWords)

	wordSubSet := make([]string, wordSubSetLen)

	for i := 0; i < wordSubSetLen; i++ {
		if (i % 5) == 0 {
			wordSubSet[i] = bagOfWords[r.Intn(len(bagOfWords))]
		} else {
			wordSubSet[i] = MakeWord(10)
		}
	}

	return bagOfWords, wordSubSet
}

func BenchmarkBuiltInMatchingSizedSets(b *testing.B) {
	largeSet, smallSet := GetTwoMatchingSizedSets(100000)

	largeSetMap := make(map[string]int, 100000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i, word := range largeSet {
			largeSetMap[word] = i
		}

		newSet := make([]string, 0, len(smallSet))

		for _, word := range smallSet {
			if _, ok := largeSetMap[word]; ok {
				newSet = append(newSet, word)
			}
		}
	}
}

func BenchmarkFastHashmapMatchingSizedSets(b *testing.B) {
	largeSet, smallSet := GetTwoMatchingSizedSets(100000)

	largeSetMap := New(100000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range largeSet {
			largeSetMap.Set(word, word)
		}

		newSet := make([]string, 0, len(smallSet))

		for _, word := range smallSet {
			if _, ok := largeSetMap.Get(word); ok {
				newSet = append(newSet, word)
			}
		}
	}
}
