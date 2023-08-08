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

	router.Static("static", "./static")

	router.GET("/", h.indexPage)

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.signUp)
		auth.POST("/sign-up", h.sendOTP)

		auth.POST("/verify/:email", h.verifyEmail)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/create-account", h.createAccount)
	}

	return router
}
