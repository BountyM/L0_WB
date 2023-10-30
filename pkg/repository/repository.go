package repository

import (
	"github.com/BountyM/L0_WB/models"
	"github.com/BountyM/L0_WB/pkg/repository/cache"
)

type Order interface {
	GetOrderByUid(uid string) (models.Order, error)
}

type Repository struct {
	Order
}

func NewRepository(cache *cache.Cache) *Repository {
	return &Repository{
		Order: NewOrderCache(cache),
	}
}
