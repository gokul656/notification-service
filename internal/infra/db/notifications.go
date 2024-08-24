package db

import "github.com/gokul656/notification-service/internal/models"

func (db *DB) InsertNewNotification() (response interface{}, err error) {
	return
}

func (db *DB) GetNotification(id string) (response interface{}, err error) {
	return
}

func (db *DB) SaveUserPreferences() (response interface{}, err error) {
	return
}

func (db *DB) GetUserPreferences(userId string) (response models.UserPreference, err error) {
	response = models.UserPreference{
		UserId:                  userId,
		MailEnabled:             true,
		SMSEnabled:              true,
		PushNotificationEnabled: true,
	}
	return
}
