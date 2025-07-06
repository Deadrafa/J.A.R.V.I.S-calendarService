package handlers

import (
	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/middleware"
	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	Name     string
	Password string
	Service  repository.CalendarService
	Metric   repository.Collector
}

func NewHandler(service repository.CalendarService, metricsCollector repository.Collector, name, pass string) *Handler {
	return &Handler{
		Service:  service,
		Metric:   metricsCollector,
		Name:     name,
		Password: pass,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	api.Use(middleware.MetricsMiddleware(h.Metric))

	api.POST("/api/create-event", h.createEventHandler)
	api.DELETE("/api/delete-event", h.deleteEventHandler)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router

}
