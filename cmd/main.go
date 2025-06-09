package main

import (
	"context"
	"flag"
	"log"

	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/config"
	"github.com/Deadrafa/J.A.R.V.I.S-calendarService/internal/handlers"
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

	handler := handlers.NewHandler(Calendar, "name", "pass")
	r := handler.InitRoutes()
	if err := r.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Panic("Cервер не поднялся: ", err)
	}

	// event := models.CalendarEvent{
	// 	Summary:     "Встреча с клиентом",
	// 	Location:    "Офис, Центральная 7",
	// 	Description: "Обсуждение нового проекта",
	// 	StartDate: models.DateTimeAndZone{
	// 		DateTime: "2025-06-09T14:00:00+07:00",
	// 		TimeZone: "Asia/Novosibirsk",
	// 	},
	// 	EndDate: models.DateTimeAndZone{
	// 		DateTime: "2025-06-09T15:30:00+07:00",
	// 		TimeZone: "Asia/Novosibirsk",
	// 	},
	// }
	// err = Calendar.CreateEvent(event)
	// if err != nil {
	// 	log.Fatalf("Ошибка CreateEvent(): %v", err)
	// }

	err = Calendar.ListEvents()
	if err != nil {
		log.Fatalf("Ошибка ListEvents(): %v", err)
	}

	// err = Calendar.DeleteEvent("m33i6oc1cauabjp0s3egaghltc")
	// if err != nil {
	// 	log.Fatalf("Ошибка DeleteEvent(): %v", err)
	// }

}
