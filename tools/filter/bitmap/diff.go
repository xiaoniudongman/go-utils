package bitmap

import (
	"github.com/xiaoniudongman/go-utils/common/consts"
	"github.com/xiaoniudongman/go-utils/tools/maths"
)

type BitMap struct {
	Bits []byte
	Max  int
}

/**
 * num/8得到byte[]的index
 * @param num
 * @return
 */
func getIndex(num int) int {
	return num >> 3
}

/**
 * num%8得到在byte[index]的位置
 * @param num
 * @return
 */
func getPosition(num int) int {
	return num & 0x07
}

/**
 * 标记指定数字（num）在bitmap中的值，标记其已经出现过<br/>
 * 将1左移position后，那个位置自然就是1，然后和以前的数据做|，这样，那个位置就替换成1了
 * @param bits
 * @param num
 */
func add(bits []byte, num int) {
	bits[getIndex(num)] |= 1 << getPosition(num)
}

/**
 * 创建bitmap数组
 */
func create(n int) []byte {
	var bits = make([]byte, getIndex(n)+1)
	for i := 0; i < n; i++ {
		add(bits, i)
	}
	return bits
}

func DifferIntsByBitMapOnBits(s []int, bitList []byte, l int) []int {
	output := make([]int, 0, len(s))
	for _, i := range s {
		index := getIndex(i)
		pos := getPosition(i)
		if bitList[index]&(1<<pos) == 0 {
			add(bitList, i)
			output = append(output, i)
		}
	}

	for _, i := range s {
		bitList[getIndex(i)] = 0
	}

	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(output))
	return output[:size:size]
}

// 长度大于30使用
func bitMapFilterInt32(s, filter []int32, bitArr BitMap, l int) []int32 {
	for _, i := range filter {
		index := i >> 3
		pos := i & 0x07
		bitArr.Bits[index] ^= 1 << pos
	}
	output := make([]int32, 0)

	for _, i := range s {
		index := i >> 3
		pos := i & 0x07
		if bitArr.Bits[index]&(1<<pos) != 0 {
			output = append(output, i)
		}
	}
	for _, i := range filter {
		index := i >> 3
		pos := i & 0x07
		bitArr.Bits[index] ^= 1 << pos
	}
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(output))
	return output[:size:size]
}

func DifferIntsByBitMap(s []int, l int) []int {
	var bitList = make([]byte, (maths.MaxInts(s)>>3)+1)
	output := make([]int, 0)
	for _, i := range s {
		index := getIndex(i)
		pos := getPosition(i)
		if bitList[index]&(1<<pos) == 0 {
			output = append(output, i)
		}
	}

	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(output))
	return output[:size:size]
}
