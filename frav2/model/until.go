package model

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var ConnectRedis *redis.Client

func UntilRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.20.21:6379", // Redis服务器地址和端口
		Password: "root",               // Redis密码，如果没有设置密码则为空
		DB:       0,                    // Redis数据库索引，默认为0
	})
	ConnectRedis = rdb
}
func SetRedis(key string, value any, mytime time.Duration) error {
	//设置键值对：使用Set方法设置键为"key"，值为"value"的数据，并设置过期时间为0
	err := ConnectRedis.Set(context.Background(), key, value, mytime).Err()
	return err
}
func GetRedis(key string) string {
	//获取键的值：使用Get方法获取键为"key"的值，并根据返回结果进行判断。
	value, err := ConnectRedis.Get(context.Background(), key).Result()
	if err == redis.Nil {
		//fmt.Println("key does not exist")
		return ""
	} else {
		//fmt.Println("key:", value)
	}
	return value
}
