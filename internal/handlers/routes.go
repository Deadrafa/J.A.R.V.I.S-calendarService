package handlers

import (
	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Name     string
	Password string
	Service  repository.CalendarService
}

func NewHandler(service repository.CalendarService, name, pass string) *Handler {
	return &Handler{
		Service:  service,
		Name:     name,
		Password: pass,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/create-event", h.createEventHandler)
	router.DELETE("/delete-event", h.deleteEventHandler)
	return router

}
