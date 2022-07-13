package rediss

import (
	"github.com/go-redis/redis"
	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/tools/errs"
)

type RedisDbInfo struct {
	redisDataDb *redis.Client
	poolSize    int
}

func createSingleClient(redisConf *config.RedisData) *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		PoolSize: redisConf.Pool_size,
		Password: redisConf.Password,
		// Database to be selected after connecting to the server.
		DB: redisConf.Db,
	})
	_, err := redisdb.Ping().Result()
	errs.CheckFatalErr(err)
	return redisdb
}

func (this *RedisDbInfo) GetRedisConnFromConf(c *config.ConfigEngine, name string) {
	redis_login := c.GetRedisDataFromConf(name)
	this.redisDataDb = createSingleClient(redis_login)
	this.poolSize = redis_login.Pool_size
}

func (this *RedisDbInfo) CreateSingleClient(addr, password string, poolSize int, db int) {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     addr,
		PoolSize: poolSize,
		Password: password,
		DB:       db,
	})
	_, err := redisDB.Ping().Result()
	errs.CheckFatalErr(err)
	this.redisDataDb = redisDB
	this.poolSize = poolSize
}

func (this *RedisDbInfo) GetDb() *redis.Client {
	return this.redisDataDb
}
