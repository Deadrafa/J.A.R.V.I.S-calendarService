package service

import (
	"context"
	"fmt"

	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/models"
	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/pkg/repository"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type Calendar struct {
	Srv    *calendar.Service
	Emails []string
}

func NewCalendar(ctx context.Context, emails []string) (repository.CalendarService, error) {

	config, err := loadOAuthConfig("internal/service/credentials.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load OAuth config: %v", err)
	}

	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to create calendar service: %v", err)
	}

	return &Calendar{
		Srv:    srv,
		Emails: emails,
	}, nil
}

func (cal *Calendar) CreateEvent(calEv models.CalendarEvent) error {
	var EventAttendees []*calendar.EventAttendee
	var FieldAttendees calendar.EventAttendee

	for _, item := range cal.Emails {
		FieldAttendees.Email = item
		EventAttendees = append(EventAttendees, &FieldAttendees)
	}
	event := calendar.Event{
		Summary:     "Встреча с клиентом",
		Location:    "Офис, 5 этаж",
		Description: "Обсуждение нового проекта",
		Start: &calendar.EventDateTime{
			DateTime: "2025-06-08T14:00:00+07:00",
			TimeZone: "Asia/Novosibirsk",
		},
		End: &calendar.EventDateTime{
			DateTime: "2025-06-08T15:30:00+07:00",
			TimeZone: "Asia/Novosibirsk",
		},
		Attendees: EventAttendees,
		Reminders: &calendar.EventReminders{
			UseDefault: false,
			Overrides: []*calendar.EventReminder{
				{Method: "email", Minutes: 24 * 60},
				{Method: "popup", Minutes: 30},
			},
			ForceSendFields: []string{"UseDefault"},
		},
	}

	createdEvent, err := cal.Srv.Events.Insert("primary", &event).Do()
	if err != nil {
		return fmt.Errorf("failed to create event: %v", err)
	}

	fmt.Printf("Event created successfully!\nView it at: %s\n", createdEvent.HtmlLink)
	return nil
}
