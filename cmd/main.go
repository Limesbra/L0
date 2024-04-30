package main

import (
	"L0/internal/cache"
	"L0/internal/database"
	"L0/internal/server"
)

func init() {
	var db database.Database
	cashe := make(cache.TypeCache)

	db.Connect()
	orders := db.GetAllOrders()
	cashe = cashe.Warming(orders)

}

func main() {
	server.RunServer()
}
