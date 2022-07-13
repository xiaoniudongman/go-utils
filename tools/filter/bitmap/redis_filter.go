package bitmap

import "github.com/xiaoniudongman/go-utils/dbs/rediss"

func UniqueByRedis(redisCli *rediss.RedisDbInfo, redisItem *rediss.ItemInfo, set []string) []string {
	_ = redisItem.ItemSetSAdd(redisCli.GetDb(), set, "asfd")
	ff, _ := redisItem.ItemGetSAdd(redisCli.GetDb(), "asfd")
	return ff
}
