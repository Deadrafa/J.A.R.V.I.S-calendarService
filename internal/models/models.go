package models

type CalendarEvent struct {
	Summary     string
	Location    string
	Description string
	StartDate   DateTimeAndZone
	EndDate     DateTimeAndZone
	Attendees   []string
}

type DateTimeAndZone struct {
	DateTime string
	TimeZone string
}
