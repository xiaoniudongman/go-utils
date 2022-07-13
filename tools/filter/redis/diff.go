package redis

import (
	"github.com/xiaoniudongman/go-utils/dbs/rediss"
	"github.com/xiaoniudongman/go-utils/tools/converter"
)

type RedisBitMap struct {
	RedisCli  *rediss.RedisDbInfo
	RedisItem *rediss.ItemInfo
}

func (r *RedisBitMap) MultiSetBit(name string, value bool, offsets ...int64) error {
	_, err := r.RedisItem.ItemPSetBit(r.RedisCli.GetDb(), offsets, converter.BoolToInt(value), name)
	return err
}

func (r *RedisBitMap) MultiSetBitStr(name string, value bool, defalut_int int64, offsets ...string) error {
	_, err := r.RedisItem.ItemPSetBit(r.RedisCli.GetDb(), converter.StrsToInt64s(offsets, defalut_int), converter.BoolToInt(value), name)
	return err
}

func (r *RedisBitMap) MultiSetBitInt(name string, value bool, offsets ...int) error {
	return r.MultiSetBit(name, value, converter.IntsToInt64s(offsets)...)
}

func (r *RedisBitMap) MultiSetBitInt32(name string, value bool, offsets ...int32) error {
	return r.MultiSetBit(name, value, converter.Int32sToInt64s(offsets)...)
}

func (r *RedisBitMap) FilterByRedis(name string, input ...int64) ([]int64, error) {
	var filter []int64
	if cmds, err := r.RedisItem.ItemPGetBit(r.RedisCli.GetDb(), input, name); err == nil {
		for i, c := range cmds {
			if !converter.Int64ToBool(c.Val()) {
				filter = append(filter, input[i])
			}
		}
	} else {
		return nil, err
	}
	return filter, nil
}

func (r *RedisBitMap) FilterByRedisInt32(name string, input ...int32) ([]int32, error) {
	ints, err := r.FilterByRedis(name, converter.Int32sToInt64s(input)...)
	return converter.Int64sToInt32s(ints), err
}

func (r *RedisBitMap) FilterByRedisInt(name string, input ...int) ([]int, error) {
	ints, err := r.FilterByRedis(name, converter.IntsToInt64s(input)...)
	return converter.Int64sToInts(ints), err
}

func (r *RedisBitMap) FilterByRedisStr(name string, defalut_int int64, input ...string) ([]string, error) {
	ints, err := r.FilterByRedis(name, converter.StrsToInt64s(input, defalut_int)...)
	return converter.Int64sToStrs(ints), err
}
