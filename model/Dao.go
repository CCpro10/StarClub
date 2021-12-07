package model

import (
	"context"
	_"fmt"
	"github.com/go-redis/redis/v8"
	_"time"
)
	//Redis相关全局变量
    var CTX = context.Background()

	var RedisDB =redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
	})

	//