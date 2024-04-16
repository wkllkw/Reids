package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	//连接user时基本不会用用户名，所以可以不填
	//密码默认为空，可以不填，所以调用db库0时可以如下所示
	//opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	opt, err := redis.ParseURL("redis://@192.168.48.112:6379/0")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	ctx := context.Background()
	val, err := rdb.Get(ctx, "v1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
