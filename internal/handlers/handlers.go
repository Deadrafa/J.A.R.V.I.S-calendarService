package handlers

import (
	"net/http"

	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createEventHandler(c *gin.Context) {

	var requestBody models.CalendarEvent

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.Service.CreateEvent(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "create",
	})
}
