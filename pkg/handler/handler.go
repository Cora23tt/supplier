package handler

import (
	"github.com/cora23tt/online-diler/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	router.Static("static", "/home/cora/Documents/projects/golang-projects/online-diler/static")

	router.GET("/", h.indexPage)

	auth := router.Group("/auth")
	{
		auth.GET("/signup", h.signUp)
		auth.POST("/signup", h.sendOTP)
		auth.POST("/verify/:email", h.verify)
		// auth.POST("/create-account", h.createAccount)
		// auth.POST("/signin", h.signIn)
	}

	return router
}
