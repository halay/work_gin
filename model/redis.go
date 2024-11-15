package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"work_gin/utils"
)

var (
	RedisDb *redis.Client
)

func InitRedis() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     utils.RedisArr,
		Password: utils.RedisPassword,
		DB:       utils.RedisDb,
		OnConnect: func(conn *redis.Conn) error {
			fmt.Println("redis连接成功")
			return nil
		},
	})
	_, err := RedisDb.Ping().Result()
	if err != nil {
		fmt.Println("连接redis失败：", err)
		os.Exit(1)
	}

}
