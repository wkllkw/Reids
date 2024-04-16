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
	val, err := rdb.Do(ctx, "get", "v1").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(val)
}
