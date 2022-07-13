package main

import (
	"fmt"
	"time"

	"github.com/xiaoniudongman/go-utils/tools/errs"

	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/dbs/rediss"
)

const (
	Config_path = "config/test.yaml"
)

func main() {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	redisItem := new(rediss.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	redisdb := new(rediss.RedisDbInfo)
	redisdb.GetRedisConnFromConf(&c, "Redis")

	// redis hset add expire
	err = redisItem.ItemHSet(redisdb.GetDb(), "ugender", "just test", "Hset:vvvv")
	if err != nil {
		fmt.Println(err)
	}
	cmd, err := redisItem.ItemHGet(redisdb.GetDb(), "ugender", "Hset:vvvv")
	if err != nil {
		fmt.Println(err)
	}
	if cmd.Val() != "just test" {
		fmt.Println(cmd.Val())
		fmt.Println("set redis item in redis db failed")
	} else {
		fmt.Println("success!!!")
	}

	time.Sleep(time.Duration(2) * time.Second)

	cmd, err = redisItem.ItemHGet(redisdb.GetDb(), "ugender", "Hset:vvvv")
	if err != nil {
		fmt.Println(err)
	}
	if cmd.Val() != "just test" {
		fmt.Println(cmd.Val())
		fmt.Println("set redis item in redis db failed")
	} else {
		fmt.Println("success!!!")
	}

	maptest := make(map[string]string, 0)
	maptest["Hset:vvvv11111111"] = "just test"
	// redis hset add expire
	_, err = redisItem.ItemPHSet(redisdb.GetDb(), "ugender", maptest)
	if err != nil {
		fmt.Println(err)
	}
	cmd, err = redisItem.ItemHGet(redisdb.GetDb(), "ugender", "Hset:vvvv11111111")
	if err != nil {
		fmt.Println(err)
	}
	if cmd.Val() != "just test" {
		fmt.Println(cmd.Val())
		fmt.Println("get1 redis item in redis db failed")
	} else {
		fmt.Println("success1!!!")
	}

	time.Sleep(time.Duration(2) * time.Second)

	cmd, err = redisItem.ItemHGet(redisdb.GetDb(), "ugender", "Hset:vvvv11111111")
	if err != nil {
		fmt.Println(err)
	}
}
