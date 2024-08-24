package repository

import "github.com/gokul656/notification-service/internal/models"

type NotificationServiceInterface interface {
	Send(recipient NotificationPayloadInterface) (response interface{}, err error)
	Broadcast(recipients []string) (response interface{}, err error)
	BroadcastToGroup(recipients []string) (response interface{}, err error)
}

type NotificationPayloadInterface interface {
	GetEventType() models.Event
	GetSubject() string
	GetDescription() string
	GetAttributes() map[string]string

	FormatTemplate(template string) (string, error)
}
