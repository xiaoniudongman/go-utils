package converter_test

import (
	"fmt"
	"testing"

	"github.com/xiaoniudongman/go-utils/tools/converter"
)

func TestStrToFloat32(b *testing.T) {
	a := "1.2"
	fmt.Println(converter.StrToFloat32(a, 9))
}

func TestStrToFloat64(b *testing.T) {
	a := "1.2"
	fmt.Println(converter.StrToFloat64(a, 9))
}
