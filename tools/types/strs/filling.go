package strs

import (
	"bytes"

	"github.com/xiaoniudongman/go-utils/common/consts"
)

// 空字符串填充
// Left justifies the text to the left
func LeftStr(text string, size int) string {
	spaces := size - len(text)
	if spaces <= 0 {
		return text
	}
	var buffer bytes.Buffer
	buffer.WriteString(text)
	for i := 0; i < spaces; i++ {
		buffer.WriteString(consts.BLANK)
	}
	return buffer.String()
}

// Right justifies the text to the right
func RightStr(text string, size int) string {
	spaces := size - len(text)
	if spaces <= 0 {
		return text
	}
	var buffer bytes.Buffer
	for i := 0; i < spaces; i++ {
		buffer.WriteString(consts.BLANK)
	}
	buffer.WriteString(text)
	return buffer.String()
}
