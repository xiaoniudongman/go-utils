package strs

import (
	"strconv"
	"testing"

	"github.com/xiaoniudongman/go-utils/tools/types/nums"

	"github.com/xiaoniudongman/go-utils/tools/maths"
)

func intsToStrs(a []int) []string {
	if nums.IsEmptyInts(a) {
		return nil
	}
	ret := make([]string, len(a))
	for i, b := range a {
		ret[i] = strconv.Itoa(b)
	}
	return ret
}

func BenchmarkBitMapDifferenceSelfInt32XXXXXXXX(b *testing.B) {
	b.ResetTimer()
	a1 := intsToStrs(maths.GenNRandInt(101, 10000))
	a2 := intsToStrs(maths.GenNRandInt(101, 10000))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DifferStrs(a1, a2)
	}
}
