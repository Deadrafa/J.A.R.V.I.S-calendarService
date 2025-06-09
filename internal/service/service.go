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
		Summary:     calEv.Summary,
		Location:    calEv.Location,
		Description: calEv.Description,
		Start: &calendar.EventDateTime{
			DateTime: calEv.StartDate.DateTime,
			TimeZone: calEv.StartDate.TimeZone,
		},
		End: &calendar.EventDateTime{
			DateTime: calEv.EndDate.DateTime,
			TimeZone: calEv.EndDate.TimeZone,
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

func (cal *Calendar) DeleteEvent(eventId string) error {
	errDelEv := cal.Srv.Events.Delete("primary", eventId).Do()
	if errDelEv != nil {
		return fmt.Errorf("failed to delete event: %v", errDelEv)
	}
	return nil
}

func (cal *Calendar) ListEvents() error {
	listEvents, err := cal.Srv.Events.List("primary").Do()
	if err != nil {
		return fmt.Errorf("failed to delete event: %v", err)
	}

	resp, err := listEvents.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to MarshalJSON(): %v", err)
	}
	fmt.Println(string(resp))

	return nil
}
