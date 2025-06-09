package models

type CalendarEvent struct {
	Summary     string          `json:"comment"`
	Location    string          `json:"location"`
	Description string          `json:"description"`
	StartDate   DateTimeAndZone `json:"start_date"`
	EndDate     DateTimeAndZone `json:"end_date"`
}

type DateTimeAndZone struct {
	DateTime string `json:"date_time"`
	TimeZone string `json:"time_zone"`
}
