package be

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	Rdb *redis.Client
)

func InitRedis() *redis.Client {
	//isTesting := os.Getenv("Testing")
	isTesting := ""
	if isTesting == "false" {
		opts, err := redis.ParseURL("redis://default:67c58cbb6e5b41bcbeba8a7a8d22266a@fly-habit.upstash.io:6379")
		if err != nil {
			panic(err)
		}
		Rdb = redis.NewClient(opts)
	} else {
		Rdb = redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
			DB:   0,
		})
	}
	err := Rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatal("err connecting tor redis, make sure docker img is running: ", err)
		return nil
	}
	return Rdb
}
