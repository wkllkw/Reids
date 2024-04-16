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
	//设置value的值
	err := rdb.Set(ctx, "v1", 11, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	//通过Result访问
	val1, err := rdb.Get(ctx, "v1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("v1", val1)

	//分别访问值和错误
	get1 := rdb.Get(ctx, "v2")
	fmt.Println(get1.Val(), get1.Err())

	//通过Result访问
	val2, err := rdb.MGet(ctx, "v1", "v2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val2)

	//分别访问值和错误
	get2 := rdb.MGet(ctx, "v1", "v2")
	fmt.Println(get2.Val(), get2.Err())
}
