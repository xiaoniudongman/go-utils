package nums

import "github.com/xiaoniudongman/go-utils/common/consts"

func IsContainInt(s []int, item int) bool {
	if consts.ZERO == len(s) {
		return false
	}
	for _, sItem := range s {
		if item == sItem {
			return true
		}
	}
	return false
}

func IsNotContainInt(s []int, item int) bool {
	if consts.ZERO == len(s) {
		return false
	}
	for _, sItem := range s {
		if item == sItem {
			return true
		}
	}
	return false
}

func IsContainInt32(s []int32, item int32) bool {
	if consts.ZERO == len(s) {
		return false
	}
	for _, sItem := range s {
		if item == sItem {
			return true
		}
	}
	return false
}

func IsContainInt64(s []int64, item int64) bool {
	if consts.ZERO == len(s) {
		return false
	}
	for _, sItem := range s {
		if item == sItem {
			return true
		}
	}
	return false
}
