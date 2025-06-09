package repository

import (
	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/models"
)

type CalendarService interface {
	CreateEvent(calEv models.CalendarEvent) error
	ListEvents() error
	DeleteEvent(eventId string) error
}
