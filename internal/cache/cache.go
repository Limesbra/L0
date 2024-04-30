package cache

import "L0/internal/model"

type TypeCache map[string]model.Order

func (TypeCache) Warming(orders map[string]model.Order) TypeCache {
	c := make(TypeCache)
	if len(orders) != 0 {
		for _, order := range orders {
			c[order.Order_uid] = order
		}
	}
	return c
}
