package redis

import (
	"fmt"
	"testing"

	"github.com/xiaoniudongman/go-utils/tools/errs"

	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/dbs/rediss"
	"github.com/xiaoniudongman/go-utils/tools/maths"
)

const length = 10000

func TestBitMapFilterByRedis(t *testing.T) {
	f := []int64{1, 2, 3, 6}
	can := []int64{1, 2, 3, 4, 5, 6}
	c := config.ConfigEngine{}
	var err error
	err = c.Load("../../config/test.yaml")
	errs.CheckCommonErr(err)
	redisItem := new(rediss.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	redisdb := new(rediss.RedisDbInfo)
	redisdb.GetRedisConnFromConf(&c, "Redis")
	kk := RedisBitMap{
		RedisCli:  redisdb,
		RedisItem: redisItem,
	}
	err = kk.MultiSetBit("ddd", true, f...)
	xx, _ := kk.FilterByRedis("ddd", can...)
	fmt.Println(xx)
}

func BenchmarkFilterByRedis(b *testing.B) {
	b.StopTimer()
	can1 := maths.GenNRandInt64(length, 100)
	can2 := maths.GenNRandInt64(length, 100)
	c := config.ConfigEngine{}
	var err error
	err = c.Load("../../config/test.yaml")
	errs.CheckCommonErr(err)
	redisItem := new(rediss.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	redisdb := new(rediss.RedisDbInfo)
	redisdb.GetRedisConnFromConf(&c, "Redis")
	kk := RedisBitMap{
		RedisCli:  redisdb,
		RedisItem: redisItem,
	}
	err = kk.MultiSetBit("ddd1", true, can1...)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = kk.FilterByRedis("ddd1", can2...)
	}
}
