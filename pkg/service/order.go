package service

import "github.com/BountyM/L0_WB/models"

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetOrderByUid(uid string) (models.Order, error) {
	return s.repo.GetOrderByUid(uid)
}
