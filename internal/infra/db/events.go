package db

import "github.com/gokul656/notification-service/internal/models"

func (db *DB) CheckConnection() bool {
	return true
}

func (db *DB) GetEventConfiguration(eventType models.Event) (response models.EventConfiguration, err error) {
	response = models.EventConfiguration{
		Type:     models.UserSignedUp,
		Priority: models.High,
	}
	return
}

func (db *DB) InsertNewEvent() (response interface{}, err error) {
	return
}

func (db *DB) GetEvent(id string) (response interface{}, err error) {
	return
}

func (db *DB) GetProtocolsByEvent(event models.Event) (response []models.EventProtocol, err error) {
	response = []models.EventProtocol{
		{
			Type:     models.UserSignedUp,
			Protocol: models.Mail,
		},
		{
			Type:     models.UserSignedUp,
			Protocol: models.SMS,
		},
	}
	return
}
