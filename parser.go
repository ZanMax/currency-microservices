package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func main() {
	fmt.Println("Parser")
	cur := apilayer_exchangerates()
	fmt.Println(cur)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)

	err = rdb.Set(ctx, "currency", cur, 0).Err()
	checkErr(err)

	val, err := rdb.Get(ctx, "currency").Result()
	checkErr(err)
	fmt.Println("currency", val)
}
