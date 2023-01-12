package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ZanMax/currency-microservices/provides"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
)

var ctx = context.Background()

func main() {
	fmt.Println("Parser")
	cur := provides.ApilayerExchangeRates()
	fmt.Println(cur)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)

	err = rdb.Set(ctx, "currency", cur, 0).Err()
	provides.CheckErr(err)

	val, err := rdb.Get(ctx, "currency").Result()
	provides.CheckErr(err)
	fmt.Println("currency", val)

	db, err := sql.Open("mysql", "root:docker@tcp(10.0.0.200:3306)/currency")
	provides.CheckErr(err)
	defer db.Close()

	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	provides.CheckErr(err2)
	fmt.Println(version)
}
