package repository

import (
	"github.com/BountyM/L0_WB/models"
	"github.com/BountyM/L0_WB/pkg/repository/cache"
)

type OrderCache struct {
	cache *cache.Cache
}

func NewOrderCache(cache *cache.Cache) *OrderCache {
	return &OrderCache{cache: cache}
}

func (r *OrderCache) GetOrderByUid(uid string) (models.Order, error) {
	var order models.Order

	order, err := r.cache.GetOrderByUid(uid)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
