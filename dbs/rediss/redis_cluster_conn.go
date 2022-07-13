package rediss

import (
	"github.com/go-redis/redis"
	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/tools/errs"
)

type RedisConnMethod interface {
	GetRedisConnFromConf(c *config.ConfigEngine, name string)
	CreateSingleClient(addr, password string, poolSize int)
}

/*
有关redis cluster连接的封装
*/
type RedisClusterDbInfo struct {
	RedisDataDb *redis.ClusterClient
	PoolSize    int
}

func getRedisNodes(redisConf *config.RedisClusterData, i int) (clusterNodes []redis.ClusterNode) {
	clusterNodes = append(clusterNodes, redis.ClusterNode{
		Addr: redisConf.Master_host[i] + ":" + redisConf.Master_port[i]})
	clusterNodes = append(clusterNodes, redis.ClusterNode{
		Addr: redisConf.Slave_host[i] + ":" + redisConf.Slave_port[i]})
	return clusterNodes
}

func readRedisCluster(redisConf *config.RedisClusterData) (redisClusters [][]redis.ClusterNode) {
	for i := 0; i < redisConf.Nodes; i++ {
		redisClusters = append(redisClusters, getRedisNodes(redisConf, i))
	}
	return redisClusters
}

func getRedisSlots(redisClusters [][]redis.ClusterNode) func() ([]redis.ClusterSlot, error) {
	return func() ([]redis.ClusterSlot, error) {
		var slots []redis.ClusterSlot
		nodes := len(redisClusters)
		var gap int
		gap = 16383 / nodes
		for i, redisClusterNodes := range redisClusters {
			slots = append(slots, redis.ClusterSlot{
				Start: i * gap,
				End:   16383 - (nodes-i-1)*gap,
				Nodes: redisClusterNodes,
			})
		}
		return slots, nil
	}
}

func createClusterClient(redisConf *config.RedisClusterData) *redis.ClusterClient {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots: getRedisSlots(readRedisCluster(redisConf)),
		PoolSize:     redisConf.Pool_size,
		Password:     redisConf.Password,
	})
	_, err := redisdb.Ping().Result()
	errs.CheckFatalErr(err)
	return redisdb
}

func (this *RedisClusterDbInfo) GetRedisConnFromConf(c *config.ConfigEngine, name string) {
	redis_login := c.GetRedisClusterDataFromConf(name)
	this.RedisDataDb = createClusterClient(redis_login)
	this.PoolSize = redis_login.Pool_size
}
