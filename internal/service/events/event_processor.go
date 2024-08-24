package events

import (
	"log"

	"github.com/gokul656/notification-service/internal/models"
	"github.com/gokul656/notification-service/internal/service/repository"
)

func (es *EventManagementService) ProcessIncomingEvent(cmd *models.IncomingEvent) {
	// check supported notification protocols for given event
	// get user preferences for type of notifications
	// check for notification priority (i.e) whether it can be ignored
	// check for schedules

	event, err := es.db.GetEventConfiguration(cmd.EventType)
	if err != nil {
		log.Printf("error: %v", err)
	}

	protocols, err := es.db.GetProtocolsByEvent(cmd.EventType)
	if err != nil {
		log.Printf("error: %v", err)
	}

	services := []repository.NotificationServiceInterface{}

	if event.Priority == models.Low {
		userPreference, err := es.db.GetUserPreferences(cmd.To)
		if err != nil {
			log.Printf("error: %v", err)
		}
		services = es.getEnabledServices(userPreference)
	} else {
		for _, protocol := range protocols {
			service := es.getNotificationServie(protocol.Protocol)
			if service != nil {
				services = append(services, service)
			}
		}
	}

	payload := models.SMSPayload{
		Attributes: map[string]string{
			"mail": cmd.To,
		},
	}

	for _, service := range services {
		_, err = service.Send(payload)
		if err != nil {
			log.Printf("err sending: %v", err)
		}
	}
}

func (es *EventManagementService) getNotificationServie(protocol models.Protocol) repository.NotificationServiceInterface {
	switch protocol {
	case models.Mail:
		return es.MailService
	case models.SMS:
		return es.SMSService
	case models.PushNotification:
		return es.PushNotificationService
	default:
		return nil
	}
}

func (es *EventManagementService) getEnabledServices(preference models.UserPreference) []repository.NotificationServiceInterface {
	services := []repository.NotificationServiceInterface{}

	if preference.MailEnabled {
		services = append(services, es.getNotificationServie(models.Mail))
	}

	if preference.SMSEnabled {
		services = append(services, es.getNotificationServie(models.SMS))
	}

	if preference.PushNotificationEnabled {
		services = append(services, es.getNotificationServie(models.PushNotification))
	}

	return services
}
