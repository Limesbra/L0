package main

import (
	"L0/internal/cache"
	"L0/internal/database"
	"L0/internal/nats"
	"L0/internal/server"
	"fmt"
)

// Инициализация и настройка базы данных, кэша и подписки на сообщения.
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

// Точка входа
func main() {
	server.RunServer()
}

// DbSubscribe подписывает базу данных на получение сообщений от NATS
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

// cacheSubscribe подписывает кэш на получение сообщений от NATS
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

// nats-streaming-server --config /Users/limesbra/wbtech/L0/internal/nats/nats_streaming.json
