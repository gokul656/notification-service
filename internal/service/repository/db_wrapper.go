package repository

import (
	"github.com/gokul656/notification-service/internal/models"
)

type DBWrapper interface {
	// health
	CheckConnection() bool

	// events
	InsertNewEvent() (interface{}, error)
	GetEvent(id string) (interface{}, error)
	GetEventConfiguration(eventType models.Event) (models.EventConfiguration, error)
	GetProtocolsByEvent(event models.Event) ([]models.EventProtocol, error)

	// notifications
	InsertNewNotification() (interface{}, error)
	GetNotification(id string) (interface{}, error)
	SaveUserPreferences() (interface{}, error)
	GetUserPreferences(userId string) (models.UserPreference, error)

	// templates
	GetMailTemplate(eventType models.Event) (string, error)
	GetSMSTemplate(eventType models.Event) (string, error)
	GetPushNotificationTemplate(eventType models.Event) (string, error)
}
