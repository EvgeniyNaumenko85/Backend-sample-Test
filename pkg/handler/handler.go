package handler

import (
	"BST/pkg/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "BST/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	//router := gin.New()
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", Ping)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	users := router.Group("/users", h.mwUserAuth)
	{
		users.GET("/current", h.getUser)
		users.GET("/", h.getAllUsers)
		users.PUT("/", h.updateUser)
		users.DELETE("/", h.deleteUser)
	}

	messages := router.Group("/messages", h.mwUserAuth, mwGetID)
	{
		messages.POST("/", h.addMessage)
		messages.GET("/:id", h.getAllMessages) //в id передается номер страницы
		messages.PUT("/:id", h.updateMessage)
		messages.DELETE("/:id", h.deleteMessage)

	}

	return router
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"reason": "up and working",
	})
}
