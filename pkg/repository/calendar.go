package repository

import (
	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/models"
)

type CalendarService interface {
	CreateEvent(calEv models.CalendarEvent) error
	// ListEvents(ctx context.Context, calID string, limit int) ([]models.CalendarEvent, error)
}
