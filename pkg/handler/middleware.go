package handler

import (
	"BST/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func mwGetID(c *gin.Context) {
	idStr, ok := c.Params.Get("id")
	if !ok {
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	c.Set("id", id)
}

func (h *Handler) mwUserAuth(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Unauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "invalid auth header"})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "token is empty"})
		return
	}

	id, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		models.ReplyError(c, err)
		c.Abort()
		return
	}
	c.Set("user_id", id)
}
