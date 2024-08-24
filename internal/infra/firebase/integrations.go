package firebase

import (
	"log"

	"github.com/gokul656/notification-service/internal/service/repository"
)

func (s *FirebaseService) Send(request repository.NotificationPayloadInterface) (response interface{}, err error) {
	template, err := s.db.GetPushNotificationTemplate(request.GetEventType())
	if err != nil {
		return nil, err
	}

	formattedPushNotification, err := request.FormatTemplate(template)
	if err != nil {
		return nil, err
	}

	log.Println("firebase:", formattedPushNotification)
	log.Println("firebase: sent push notification")
	return
}

func (s *FirebaseService) Broadcast(recipients []string) (response interface{}, err error) {
	log.Println("firebase: sent broadcast notification")
	return
}

func (s *FirebaseService) BroadcastToGroup(recipients []string) (response interface{}, err error) {
	log.Println("firebase: sent group notification")
	return
}
