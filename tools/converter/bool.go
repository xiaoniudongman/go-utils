package converter

import "github.com/xiaoniudongman/go-utils/common/consts"

/**
int 转换
*/
func BoolToInt(b bool) int {
	if true == b {
		return consts.ONE
	}
	return consts.ZERO
}

func BoolToInt32(b bool) int32 {
	return IntToInt32(BoolToInt(b))
}

func BoolToInt64(b bool) int64 {
	return IntToInt64(BoolToInt(b))
}

func Int64ToBool(b int64) bool {
	if b > 0 {
		return true
	}
	return false
}
