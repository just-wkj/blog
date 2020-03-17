package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

var redisdb *redis.Client
func initRedisDb()(err error){
	redisdb = redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		Password:"",
		DB:0,
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		return
	}
	return
}

func main() {
	err := initRedisDb()
	if err != nil {
		fmt.Println("redis 连接失败")
		return
	}
	fmt.Println("redis ok")
	//<<redis实战>>


	// zset
	key := "rank"
	items := []*redis.Z{
		&redis.Z{Score:88, Member:"php"},
		&redis.Z{Score:86, Member:"python"},
		&redis.Z{Score:91, Member:"go"},
		&redis.Z{Score:84, Member:"java"},
	}
	//ZRANGE rank 0 2
	redisdb.ZAdd(key, items...)

	//给java加分
	score, err := redisdb.ZIncrBy(key, 10, "java").Result()
	if err != nil {
		fmt.Println("修改错误!")
		return
	}
	fmt.Println("score:%d", score)
}
