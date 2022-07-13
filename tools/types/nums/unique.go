package nums

import (
	"github.com/xiaoniudongman/go-utils/common/consts"
)

var PerfThrUnique = 300

// int slice自去重,通过两重循环过滤重复元素
func uniqueIntByLoop(s []int) []int {
	var dup []int
	for _, v := range s {
		if !IsContainInt(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueIntByMap(s []int) []int {
	var dup []int
	tmpMap := make(map[int]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = consts.ZERO
			dup = append(dup, v)
		}
	}
	return dup
}

// int slice自去重,通过两重循环过滤重复元素
func uniqueIntByLoop32(s []int32) []int32 {
	var dup []int32
	for _, v := range s {
		if !IsContainInt32(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueIntByMap32(s []int32) []int32 {
	var dup []int32
	tmpMap := make(map[int32]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = consts.ZERO
			dup = append(dup, v)
		}
	}
	return dup
}

// int slice自去重,通过两重循环过滤重复元素
func uniqueIntByLoop64(s []int64) []int64 {
	var dup []int64
	for _, v := range s {
		if !IsContainInt64(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueIntByMap64(s []int64) []int64 {
	var dup []int64
	var tmpMap = make(map[int64]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = consts.ZERO
			dup = append(dup, v)
		}
	}
	return dup
}

// list自去重
func UniqueInt(s []int) []int {
	if len(s) < PerfThrUnique {
		return uniqueIntByLoop(s)
	}
	return uniqueIntByMap(s)
}

func UniqueInt32(s []int32) []int32 {
	if len(s) < PerfThrUnique {
		return uniqueIntByLoop32(s)
	}
	return uniqueIntByMap32(s)
}

func UniqueIntLen(s []int, l int) []int {
	return CutIntList(UniqueInt(s), l)
}

func UniqueInt32Len(s []int32, l int) []int32 {
	return CutInt32List(UniqueInt32(s), l)
}
