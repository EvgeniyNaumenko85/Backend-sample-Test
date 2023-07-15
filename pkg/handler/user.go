package handler

import (
	"BST/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Current User
// @Security BearerAuth
// @SecurityScheme BearerAuth bearer Authorization header using the Bearer scheme
// @Tags users
// @Description get current user
// @ID get-current-user
// @Produce  json
// @Param Authorization header string true "Bearer Token Authorization" default(Bearer <ваш JWT-токен>)
// @Success 200 {string} integer
// @Failure 400,404 {string} models.BadRequest
// @Failure 401 {string} models.Unauthorized
// @Failure 500 {string} models.ReplyError
// @Failure default {} string "internal Server Error"
// @Router /users/current [get]
func (h *Handler) getUser(c *gin.Context) {
	userID := c.GetInt("user_id")
	user, err := h.services.GetUser(c.Request.Context(), userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Get All Users
// @Security BearerAuth
// @SecurityScheme BearerAuth bearer Authorization header using the Bearer scheme
// @Tags users
// @Description get all users
// @ID get-all-user
// @Produce  json
// @Param Authorization header string true "Bearer Token Authorization" default(Bearer <ваш JWT-токен>)
// @Success 200 {string} integer
// @Failure 400,404 {string} models.BadRequest
// @Failure 401 {string} models.Unauthorized
// @Failure 500 {string} models.ReplyError
// @Failure default {} string "internal Server Error"
// @Router /users/ [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.GetAllUsers(c.Request.Context())
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Update User (self)
// @Security BearerAuth
// @SecurityScheme BearerAuth bearer Authorization header using the Bearer scheme
// @Tags users
// @Description update user (self)
// @ID update-user-(self)
// @Produce  json
// @Param Authorization header string true "Bearer Token Authorization" default(Bearer <ваш JWT-токен>)
// @Param input body models.RequestUser false "Input the message body to update"
// @Success 200 {string} integer
// @Failure 400,404 {string} models.BadRequest
// @Failure 401 {string} models.Unauthorized
// @Failure 500 {string} models.ReplyError
// @Failure default {} string "internal Server Error"
// @Router /users/ [put]
func (h *Handler) updateUser(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	u.ID = c.GetInt("user_id")
	err := h.services.UpdateUser(c.Request.Context(), u)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

// @Summary Delete User (self)
// @Security BearerAuth
// @SecurityScheme BearerAuth bearer Authorization header using the Bearer scheme
// @Tags users
// @Description delete user (self)
// @ID delete-user-(self)
// @Produce  json
// @Param Authorization header string true "Bearer Token Authorization" default(Bearer <ваш JWT-токен>)
// @Success 200 {string} integer
// @Failure 400,404 {string} models.BadRequest
// @Failure 401 {string} models.Unauthorized
// @Failure 500 {string} models.ReplyError
// @Failure default {} string "internal Server Error"
// @Router /users/ [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userID := c.GetInt("user_id")
	err := h.services.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
