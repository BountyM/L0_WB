package handler

import (
	"github.com/BountyM/L0_WB/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	order := router.Group("/order")
	{
		order.GET("/:uid", h.getOrderByUid)
	}

	return router
}
