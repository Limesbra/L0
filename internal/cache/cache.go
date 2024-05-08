package cache

import (
	"L0/internal/model"
	"encoding/json"
	"fmt"
)

// TypeCache представляет кэш заказов.
type TypeCache map[string]model.Order

// PtrCache - указатель на TypeCache.
var PtrCache *TypeCache

// Warming загружает данные заказов в кэш.
func (cache TypeCache) Warming(orders map[string]model.Order) TypeCache {
	c := make(TypeCache)
	if len(orders) != 0 {
		for _, order := range orders {
			c[order.Order_uid] = order
		}
	}
	return c
}

// AddData добавляет заказ в кэш.
func (cache TypeCache) AddData(order model.Order) {
	cache[order.Order_uid] = order
}

// Consume обрабатывает данные и добавляет заказ в кэш.
func (cache TypeCache) Consume(data []byte) error {
	var order model.Order
	err := json.Unmarshal(data, &order)
	if err != nil {
		fmt.Println("There", err)
		return err
	}
	cache.AddData(order)
	return nil
}
