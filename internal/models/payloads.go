package models

import (
	"fmt"
	"strings"
)

type SMSPayload struct {
	EventType   Event             `json:"event"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

func (p SMSPayload) GetEventType() Event              { return p.EventType }
func (p SMSPayload) GetSubject() string               { return p.Title }
func (p SMSPayload) GetDescription() string           { return p.Description }
func (p SMSPayload) GetAttributes() map[string]string { return p.Attributes }
func (p SMSPayload) FormatTemplate(template string) (string, error) {
	for key, value := range p.Attributes {
		template = strings.ReplaceAll(template, fmt.Sprintf("{{%s}}", key), value)
	}
	return template, nil
}
