package handler

import (
	"BST/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.RequestUser true "account info"
// @Success 200 {string} models.OK
// @Failure 400,404 {string} models.BadRequest
// @Failure 406 {string} models.ErrDuplicate
// @Failure 500 {string} models.InternalErr
// @Failure default {} string "internal Server Error"
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	err := h.services.AddUser(c.Request.Context(), u)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.RequestUser true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {string} models.BadRequest
// @Failure 500 {string} models.InternalErr
// @Failure default {} string "internal Server Error"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	token, err := h.services.GenerateToken(c.Request.Context(), u)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
