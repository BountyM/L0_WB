package cache

import (
	"errors"
	"sync"

	"github.com/BountyM/L0_WB/models"
)

type Cache struct {
	sync.RWMutex
	orders map[string]models.Order
}

func NewCache() *Cache {
	return &Cache{
		orders: make(map[string]models.Order),
	}
}

func (c *Cache) Set(key string, Order models.Order) {
	c.Lock()
	defer c.Unlock()
	c.orders[key] = Order
}

func (c *Cache) GetOrderByUid(key string) (models.Order, error) {
	c.RLock()
	defer c.RUnlock()
	order, ok := c.orders[key]
	if !ok {
		return order, errors.New("error cach GetOrderByUid")
	}
	return order, nil
}
