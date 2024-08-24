package main

import (
	"github.com/gokul656/notification-service/internal/infra/db"
	"github.com/gokul656/notification-service/internal/infra/firebase"
	"github.com/gokul656/notification-service/internal/infra/mail"
	"github.com/gokul656/notification-service/internal/infra/sms"
	"github.com/gokul656/notification-service/internal/models"
	"github.com/gokul656/notification-service/internal/service/events"
)

func main() {
	db := db.New()
	smsService := sms.New(db)
	mailService := mail.New(db)
	pushService := firebase.New(db)

	eventManagementService := events.New(db, smsService, mailService, pushService)

	event := &models.IncomingEvent{
		From:         "noreply@codes.com",
		To:           "gokul656@gmail.com",
		EventType:    models.UserSignedUp,
		Participents: []string{},
		IsScheduled:  false,
		ScheduledAt:  nil,
	}
	eventManagementService.ProcessIncomingEvent(event)

	event.EventType = models.UserDeposited
	eventManagementService.ProcessIncomingEvent(event)
}
