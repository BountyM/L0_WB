package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserByUid(c *gin.Context) {
	uid := c.Params.ByName("uid")

	order, err := h.services.Order.GetUserByUid(uid)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}
