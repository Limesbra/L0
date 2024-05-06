package cache

import (
	"L0/internal/model"
	"encoding/json"
	"fmt"
)

type TypeCache map[string]model.Order

var PtrCache *TypeCache

func (cache TypeCache) Warming(orders map[string]model.Order) TypeCache {
	c := make(TypeCache)
	if len(orders) != 0 {
		for _, order := range orders {
			c[order.Order_uid] = order
		}
	}
	return c
}

func (cache TypeCache) AddData(order model.Order) {
	cache[order.Order_uid] = order
}

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
