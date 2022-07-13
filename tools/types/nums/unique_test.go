package nums

import (
	"fmt"
	"testing"

	"github.com/xiaoniudongman/go-utils/tools/maths"
)

const length = 30000

func TestRmDuplicateIntLen(b *testing.T) {
	a := []int{1, 2, 3, 4, 5, 4, 4, 6}
	fmt.Println(UniqueIntLen(a, 9))
}

func BenchmarkUniqueIntByLoop(b *testing.B) {
	b.StopTimer()
	a := maths.GenNRandInt(length, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		uniqueIntByLoop(a)
	}
}

func BenchmarkUniqueIntByMap(b *testing.B) {
	b.StopTimer()
	a := maths.GenNRandInt(length, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		uniqueIntByMap(a)
	}
}

func BenchmarkUniqueInt(b *testing.B) {
	b.ResetTimer()
	a := maths.GenNRandInt(length, 100)
	for i := 0; i < b.N; i++ {
		UniqueInt(a)
	}
}

func BenchmarkDifferInt(b *testing.B) {
	a := maths.GenNRandInt(length, 100)
	aa := maths.GenNRandInt(length, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DifferInt(a, aa)
	}
}

func BenchmarkDifferIntByMap(b *testing.B) {
	a := maths.GenNRandInt(length, 100)
	aa := maths.GenNRandInt(length, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DifferInt(a, aa)
	}
}
