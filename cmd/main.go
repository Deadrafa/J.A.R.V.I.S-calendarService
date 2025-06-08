package main

import (
	"context"
	"flag"
	"log"

	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/config"
	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/models"
	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/service"
)

func main() {
	cfgPath := flag.String("cfg", "internal/config/default.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(*cfgPath)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}

	Calendar, err := service.NewCalendar(context.Background(), cfg.Emails)
	if err != nil {
		log.Fatalf("Ошибка NewCalendar(): %v", err)
	}
	err = Calendar.CreateEvent(models.CalendarEvent{})
	if err != nil {
		log.Fatalf("Ошибка CreateEvent(): %v", err)
	}

}
