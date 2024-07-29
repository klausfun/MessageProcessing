package handler

import (
	"MessageProcessing/pkg/service"
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

	api := router.Group("/api")
	{
		message := api.Group("/message")
		{
			message.POST("/", h.createMessage)
			message.GET("/", h.getAllMessages)
		}
	}

	return router
}
