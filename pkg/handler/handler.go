package handler

import "github.com/gin-gonic/gin"

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
		order.GET("/:uid", h.getUserByUid)
	}

	return router
}
