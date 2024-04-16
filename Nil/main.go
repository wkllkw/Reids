package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.48.112:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	val, err := rdb.Get(ctx, "key").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key不存在")
	case err != nil:
		fmt.Println("错误", err)
	case val == "":
		fmt.Println("值是空字符串")
	}

}
