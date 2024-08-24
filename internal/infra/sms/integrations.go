package sms

import (
	"log"

	"github.com/gokul656/notification-service/internal/service/repository"
)

func (s *SMSService) Send(request repository.NotificationPayloadInterface) (response interface{}, err error) {
	template, err := s.db.GetSMSTemplate(request.GetEventType())
	if err != nil {
		return nil, err
	}

	formattedSMS, err := request.FormatTemplate(template)
	if err != nil {
		return nil, err
	}

	log.Println("smsservice:", formattedSMS)
	log.Println("smsservice: sent push notification")
	return
}

func (s *SMSService) Broadcast(recipients []string) (response interface{}, err error) {
	log.Println("smsservice: sent broadcast notification")
	return

}

func (s *SMSService) BroadcastToGroup(recipients []string) (response interface{}, err error) {
	log.Println("smsservice: sent group notification")
	return
}
