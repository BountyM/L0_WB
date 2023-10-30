package service

import (
	"github.com/BountyM/L0_WB/models"
	"github.com/BountyM/L0_WB/pkg/repository"
)

type Order interface {
	GetOrderByUid(uid string) (models.Order, error)
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
