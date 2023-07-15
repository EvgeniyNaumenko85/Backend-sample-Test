package handler

import (
	"net/http"

	"BST/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create Message
// @Security BearerAuth
// @SecurityScheme BearerAuth bearer Authorization header using the Bearer scheme
// @Tags messages
// @Description create message
// @ID create-message
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token Authorization" default(Bearer <ваш JWT-токен>)
// @Param input body models.RequestMessage true "Input the message body"
// @Success 200 {string} integer
// @Failure 400,404 {string} models.BadRequest
// @Failure 401 {string} models.Unauthorized
// @Failure 500 {string} models.ReplyError
// @Failure default {} string "internal Server Error"
// @Router /messages [post]
func (h *Handler) addMessage(c *gin.Context) {
	var t models.Message
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	if len(t.Description) == 0 || len(t.Description) > 256 {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.BadRequest)
		return
	}

	t.UserID = c.GetInt("user_id")

	id, err := h.services.AddMessage(c.Request.Context(), t)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

// @Summary Delete Message By ID
// @Security BearerAuth
// @SecurityScheme BearerAuth bearer Authorization header using the Bearer scheme
// @Tags messages
// @Description  Delete message by ID
// @ID delete-message-by-ID
// @Produce  json
// @Param Authorization header string true "Bearer Token Authorization" default(Bearer <ваш JWT-токен>)
// @Param id path int true "ID of the message"
// @Success 200 {string} integer
// @Failure 400,404 {string} models.BadRequest
// @Failure 401 {string} models.Unauthorized
// @Failure 500 {string} models.ReplyError
// @Failure default {} string "internal Server Error"
// @Router /messages/{id} [delete]
func (h *Handler) deleteMessage(c *gin.Context) {
	id := c.GetInt("id")
	userID := c.GetInt("user_id")
	err := h.services.DeleteMessage(c.Request.Context(), id, userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

// @Summary Update Message By ID
// @Security BearerAuth
// @SecurityScheme BearerAuth bearer Authorization header using the Bearer scheme
// @Tags messages
// @Description update message by ID
// @ID update-message-by-ID
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token Authorization" default(Bearer <ваш JWT-токен>)
// @Param id path int true "ID of the message"
// @Param input body models.RequestMessage false "Input the message body to update"
// @Success 200 {string} integer
// @Failure 400,404 {string} models.BadRequest
// @Failure 401 {string} models.Unauthorized
// @Failure 500 {string} models.ReplyError
// @Failure default {} string "internal Server Error"
// @Router /messages/{id} [put]
func (h *Handler) updateMessage(c *gin.Context) {
	var t models.Message
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	userID := c.GetInt("user_id")
	id := c.GetInt("id")

	err := h.services.UpdateMessage(c.Request.Context(), id, userID, t)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

// @Summary Get All Message by PageID
// @Security BearerAuth
// @SecurityScheme BearerAuth bearer Authorization header using the Bearer scheme
// @Tags messages
// @Description get all message by pageID
// @ID get-all-message-by-pageID
// @Produce  json
// @Param Authorization header string true "Bearer Token Authorization" default(Bearer <ваш JWT-токен>)
// @Param id path int true "Input page ID to pagination"
// @Success 200 {string} integer
// @Failure 400,404 {string} models.BadRequest
// @Failure 401 {string} models.Unauthorized
// @Failure 500 {string} models.ReplyError
// @Failure default {} string "internal Server Error"
// @Router /messages/{id} [get]
func (h *Handler) getAllMessages(c *gin.Context) {
	userID := c.GetInt("user_id")
	PageID := c.GetInt("id")
	messages, err := h.services.GetAllMessages(c.Request.Context(), PageID, userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, messages)
}
