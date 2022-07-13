package nums

import "github.com/xiaoniudongman/go-utils/common/consts"

var PerfThrDiff = 100

// s1对s2做差
func differIntByLoop(s1, s2 []int) []int {
	if IsEmptyInts(s2) {
		return s1
	}
	var ints []int
	for _, i := range s1 {
		if !IsContainInt(s2, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

// 通过map主键唯一的特性过滤重复元素
func differIntByMap(s1, s2 []int) []int {
	var dup []int
	var tmpMap = make(map[int]byte, len(s2)) // 存放不重复主键
	for _, v := range s2 {
		tmpMap[v] = consts.ZERO
	}
	for _, i := range s1 {
		if _, ok := tmpMap[i]; !ok {
			dup = append(dup, i)
		}
	}
	return dup
}

func differIntByLoop32(s1, s2 []int32) []int32 {
	if IsEmptyInt32s(s2) {
		return s1
	}
	var ints []int32
	for _, i := range s1 {
		if !IsContainInt32(s2, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

// 通过map主键唯一的特性过滤重复元素
func differIntByMap32(s1, s2 []int32) []int32 {
	var dup []int32
	tmpMap := make(map[int32]byte, len(s2)) // 存放不重复主键
	for _, v := range s2 {
		tmpMap[v] = consts.ZERO
	}
	for _, i := range s1 {
		if _, ok := tmpMap[i]; !ok {
			dup = append(dup, i)
		}
	}
	return dup
}

func DifferInt(s1, s2 []int) []int {
	if len(s1) < PerfThrDiff {
		return differIntByLoop(s1, s2)
	}
	return differIntByMap(s1, s2)
}

func DifferIntLen(s1, s2 []int, l int) []int {
	return CutIntList(DifferInt(s1, s2), l)
}

func DifferInt32(s1, s2 []int32) []int32 {
	if len(s1) < PerfThrDiff {
		return differIntByLoop32(s1, s2)
	}
	return differIntByMap32(s1, s2)
}

func DifferInt32Len(s1, s2 []int32, l int) []int32 {
	return CutInt32List(DifferInt32(s1, s2), l)
}
