package main

import (
	"L0/internal/cache"
	"L0/internal/database"
	"L0/internal/nats"
	"L0/internal/server"
	"fmt"
)

func init() {
	var db database.Database
	cache_ := make(cache.TypeCache)

	db.Connect()
	orders := db.GetAllOrders()
	cache_ = cache_.Warming(orders)
	cacheSubscribe(&cache_)
	DbSubscribe(&db)
	cache.PtrCache = &cache_

}

func main() {
	server.RunServer()
}

func DbSubscribe(db *database.Database) {
	var srvNats nats.Service
	err := srvNats.Connect("consumer_db")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = srvNats.Subscribe("upload_consumer", db)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func cacheSubscribe(c *cache.TypeCache) {
	var srvNats nats.Service
	err := srvNats.Connect("consumer_cache")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = srvNats.Subscribe("upload_consumer", c)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// nats-streaming-server --config /Users/limesbra/wbtech/L0/internal/nats/another_nats.json
