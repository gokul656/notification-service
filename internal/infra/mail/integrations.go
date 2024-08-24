package mail

import (
	"log"

	"github.com/gokul656/notification-service/internal/service/repository"
)

func (s *MailService) Send(request repository.NotificationPayloadInterface) (response interface{}, err error) {

	template, err := s.db.GetMailTemplate(request.GetEventType())
	if err != nil {
		return nil, err
	}

	formattedMail, err := request.FormatTemplate(template)
	if err != nil {
		return nil, err
	}

	log.Println("mailservice:", formattedMail)
	log.Println("mailservice: sent push notification")

	return
}

func (s *MailService) Broadcast(recipients []string) (response interface{}, err error) {
	log.Println("mailservice: sent broadcast notification")

	return
}

func (s *MailService) BroadcastToGroup(recipients []string) (response interface{}, err error) {
	log.Println("mailservice: sent group notification")
	return
}
