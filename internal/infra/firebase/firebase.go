package firebase

import "github.com/gokul656/notification-service/internal/service/repository"

type FirebaseService struct {
	db repository.DBWrapper
}

func New(db repository.DBWrapper) *FirebaseService {
	return &FirebaseService{
		db: db,
	}
}
