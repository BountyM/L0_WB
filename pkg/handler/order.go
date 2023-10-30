package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getOrderByUid(c *gin.Context) {
	uid := c.Params.ByName("uid")

	order, err := h.services.Order.GetOrderByUid(uid)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}
