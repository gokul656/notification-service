package sms

import "github.com/gokul656/notification-service/internal/service/repository"

type SMSService struct {
	db repository.DBWrapper
}

func New(db repository.DBWrapper) *SMSService {
	return &SMSService{
		db: db,
	}
}
