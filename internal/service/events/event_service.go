package events

import (
	"github.com/gokul656/notification-service/internal/service/repository"
	ns "github.com/gokul656/notification-service/internal/service/repository"
)

type EventManagementService struct {
	db                      repository.DBWrapper
	SMSService              ns.NotificationServiceInterface
	MailService             ns.NotificationServiceInterface
	PushNotificationService ns.NotificationServiceInterface
}

func New(
	db repository.DBWrapper,
	smsService ns.NotificationServiceInterface,
	mailService ns.NotificationServiceInterface,
	pushNotificationService ns.NotificationServiceInterface) *EventManagementService {
	return &EventManagementService{
		db:                      db,
		SMSService:              smsService,
		MailService:             mailService,
		PushNotificationService: pushNotificationService,
	}
}
