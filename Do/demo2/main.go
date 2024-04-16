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
	// Text 是 get.Val().(string) 的快捷方式，具有正确的错误处理。
	val, err := rdb.Do(ctx, "get", "v1").Text()
	fmt.Println(val, err)
}
