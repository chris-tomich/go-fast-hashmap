package uuid

import (
	"testing"
	"math/rand"
	"time"
	"strconv"
	"fmt"
	"github.com/satori/go.uuid"
)

func TestGoFastHashmap(t *testing.T) {
	m := New(10)

	a := uuid.NewV4()
	b := uuid.NewV4()
	c := uuid.NewV4()
	d := uuid.NewV4()

	m.Set(a, 0)
	m.Set(b, 1)
	m.Set(c, 2)
	m.Set(d, 3)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

//func BenchmarkFindNextPrime(b *testing.B) {
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		b.StopTimer()
//		testNum := uint64(r.Int63n(int64(100000000)))
//		b.StartTimer()
//
//		nextPrime(testNum)
//	}
//}

func MakeWord(maxSize int) string {
	var letters = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	wordSize := r.Intn(maxSize);

	if wordSize < 30 {
		wordSize = 30
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

func GetTwoMatchingSizedSets(size int) ([]uuid.UUID, []uuid.UUID) {
	bagOfUuids := make([]uuid.UUID, 0, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		bagOfUuids = append(bagOfUuids, uuid.NewV4())
	}

	wordSubSetLen := len(bagOfUuids)

	wordSubSet := make([]uuid.UUID, wordSubSetLen)

	for i := 0; i < wordSubSetLen; i++ {
		if (i % 5) == 0 {
			wordSubSet[i] = bagOfUuids[r.Intn(len(bagOfUuids))]
		} else {
			wordSubSet[i] = uuid.NewV4()
		}
	}

	return bagOfUuids, wordSubSet
}

func BenchmarkBuiltInMatchingSizedSets(b *testing.B) {
	largeSet, smallSet := GetTwoMatchingSizedSets(100000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		largeSetMap := make(map[myHashType]int, 100000)

		for i, uid := range largeSet {
			largeSetMap[hashUuidMyHashType(uid)] = i
		}

		newSet := make([]uuid.UUID, 0, len(smallSet))

		for _, uid := range smallSet {
			if _, ok := largeSetMap[hashUuidMyHashType(uid)]; ok {
				newSet = append(newSet, uid)
			}
		}
	}
}

func BenchmarkBuiltInByteArrayMatchingSizedSets(b *testing.B) {
	largeSet, smallSet := GetTwoMatchingSizedSets(100000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		largeSetMap := make(map[uuid.UUID]int, 100000)

		for i, uid := range largeSet {
			largeSetMap[uid] = i
		}

		newSet := make([]uuid.UUID, 0, len(smallSet))

		for _, uid := range smallSet {
			if _, ok := largeSetMap[uid]; ok {
				newSet = append(newSet, uid)
			}
		}
	}
}

//func BenchmarkHash(b *testing.B) {
//	largeSet, _ := GetTwoMatchingSizedSets(100000)
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		for _, uuid := range largeSet {
//			hashUuid(uuid)
//		}
//	}
//}
//
//func BenchmarkFastHashmapMatchingSizedSets(b *testing.B) {
//	largeSet, smallSet := GetTwoMatchingSizedSets(100000)
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		largeSetMap := New(100000)
//
//		for i, uuid := range largeSet {
//			largeSetMap.Set(uuid, i)
//		}
//
//		newSet := make([]uuid.UUID, 0, len(smallSet))
//
//		for _, uuid := range smallSet {
//			if _, ok := largeSetMap.Get(uuid); ok {
//				newSet = append(newSet, uuid)
//			}
//		}
//	}
//}
