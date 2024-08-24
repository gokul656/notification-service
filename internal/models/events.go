package models

import "time"

type Event string

const (
	UserSignedUp  Event = "user_signed_up"
	UserLoggedIn  Event = "user_logged_in"
	UserDeposited Event = "user_deposited"
)

type Priority string

const (
	Low  Priority = "low"  // can be ignored if user turned off notifications
	High Priority = "high" // must be sent even if user turned off notifications
)

type Protocol string

const (
	SMS              Protocol = "sms"
	Mail             Protocol = "mail"
	PushNotification Protocol = "push_notification"
)

type IncomingEvent struct {
	From         string     `json:"from"`
	To           string     `json:"to"`
	EventType    Event      `json:"event_type"`
	Participents []string   `json:"participents"`
	IsScheduled  bool       `json:"is_scheduled"`
	ScheduledAt  *time.Time `json:"scheduled_at"`
}

type EventConfiguration struct {
	Id       string   `db:"id"`
	Type     Event    `db:"type"`
	Priority Priority `db:"priority"`
}

type EventProtocol struct {
	Id       string   `db:"id"`
	Type     Event    `db:"type"`
	Protocol Protocol `db:"protocol"`
}

type UserPreference struct {
	UserId                  string `db:"user_id"`
	MailEnabled             bool   `db:"mail_enabled"`
	SMSEnabled              bool   `db:"sms_enabled"`
	PushNotificationEnabled bool   `db:"push_notification_enabled"`
}
