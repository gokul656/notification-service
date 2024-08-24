package mail

import "github.com/gokul656/notification-service/internal/service/repository"

type MailService struct {
	db repository.DBWrapper
}

func New(db repository.DBWrapper) *MailService {
	return &MailService{
		db: db,
	}
}
