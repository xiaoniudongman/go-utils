package bitmap

import (
	"testing"

	"github.com/xiaoniudongman/go-utils/tools/maths"
)

/**
 * num/8得到byte[]的index
 * @param num
 * @return
 */
func TestBitMapDifferenceSelfInt32XXXXXXXX(t *testing.T) {
	a := maths.GenNRandInt(3000, 100)
	DifferIntsByBitMap(a, 10)
}

func BenchmarkBitMapDifferenceSelfInt32XXXXXXXX(b *testing.B) {
	b.ResetTimer()
	a := maths.GenNRandInt(30000, 100)
	//b.Log(a)
	for i := 0; i < b.N; i++ {
		DifferIntsByBitMap(a, 10)
	}
}

func BenchmarkBitMapDifferenceSelfInt32(b *testing.B) {
	b.StopTimer()
	a := maths.GenNRandInt(300, 100)
	var bitList = make([]byte, 100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DifferIntsByBitMapOnBits(a, bitList, 10)
	}
}
