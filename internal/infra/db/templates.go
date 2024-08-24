package db

import "github.com/gokul656/notification-service/internal/models"

func (db *DB) GetMailTemplate(eventType models.Event) (template string, err error) {
	return "sign up success for {{mail}} ", nil
}

func (db *DB) GetSMSTemplate(eventType models.Event) (template string, err error) {
	return "sign up success for {{mail}} ", nil
}

func (db *DB) GetPushNotificationTemplate(eventType models.Event) (template string, err error) {
	return "notification sign up success for {{mail}}", nil
}
