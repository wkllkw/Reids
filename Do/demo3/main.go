package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.48.112:6379", //ip地址，端口号
		Password: "",                    //没有密码，默认值
		DB:       0,                     //默认DB
	})

	ctx := context.Background()
	// 注意此处set的使用方法
	err := rdb.Do(ctx, "set", "v3", "33").Err()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key不存在")
			return
		}
		panic(err)
	}
	val, err := rdb.Get(ctx, "v3").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val, err)
}
