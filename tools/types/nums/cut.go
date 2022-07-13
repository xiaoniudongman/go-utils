package nums

import (
	"fmt"

	"github.com/xiaoniudongman/go-utils/common/consts"
	"github.com/xiaoniudongman/go-utils/tools/maths"
)

// cut strs
func CutIntList(s []int, l int) []int {
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(s))
	return s[:size:size]
}

func CutIntListAndFilling(s []int, f []int, l int) []int {
	if cut := CutIntList(s, l); len(cut) < l {
		tmp := append(s, f...)
		return UniqueIntLen(tmp, l)
	} else {
		return cut
	}
}

func CutInt32List(s []int32, l int) []int32 {
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(s))
	return s[:size:size]
}

func CutInt32ListAndFilling(s []int32, f []int32, l int) []int32 {
	if cut := CutInt32List(s, l); len(cut) < l {
		tmp := append(s, f...)
		return UniqueInt32Len(tmp, l)
	} else {
		return cut
	}
}

// int list Loop
func GetIntListNoLoop(s []int, size, num int) ([]int, error) {
	if num <= consts.ZERO || size <= consts.ZERO {
		return nil, fmt.Errorf("Input paras error!!!")
	}
	minThr := maths.MinInt(size*(num-1), len(s))
	maxThr := maths.MinInt(size*num, len(s))
	return s[minThr:maxThr], nil
}

func GetIntListLoop(s []int, size, num int) ([]int, error) {
	if num <= consts.ZERO || size <= consts.ZERO || size >= len(s) {
		return nil, fmt.Errorf("Input parameter error!!!")
	}
	start := (size * (num - 1)) % len(s)
	end := (num * size) % len(s)
	if start < end {
		return s[start:end:end], nil
	} else {
		out := make([]int, consts.ZERO, size)
		out = append(out, s[start:len(s):len(s)]...)
		out = append(out, s[:end:end]...)
		return out, nil
	}
}
