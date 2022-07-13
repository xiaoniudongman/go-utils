package nums

import (
	"github.com/xiaoniudongman/go-utils/common/consts"
	"github.com/xiaoniudongman/go-utils/tools/maths"
)

// ints mix
func MixListIntV2(s ...[]int) []int {
	var count int
	var maxLen int
	for _, item := range s {
		count += len(item)
		maxLen = maths.MaxInt(maxLen, len(item))
	}
	mix := make([]int, 0, count)
	for i := 0; i < maxLen; i++ {
		for j := range s {
			if i < len(s[j]) {
				mix = append(mix, s[j][i])
			}
		}
	}
	return UniqueInt(mix)
}

func MixListIntLenV2(l int, s ...[]int) []int {
	mix := MixListIntV2(s...)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(mix))
	return mix[:size:size]
}

func MixListInt32V2(s ...[]int32) []int32 {
	var count int
	var maxLen int
	for _, item := range s {
		count += len(item)
		maxLen = maths.MaxInt(maxLen, len(item))
	}
	mix := make([]int32, 0, count)
	for i := 0; i < maxLen; i++ {
		for j := range s {
			if i < len(s[j]) {
				mix = append(mix, s[j][i])
			}
		}
	}
	return UniqueInt32(mix)
}

func MixListInt32LenV2(l int, s ...[]int32) []int32 {
	mix := MixListInt32V2(s...)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(mix))
	return mix[:size:size]
}
