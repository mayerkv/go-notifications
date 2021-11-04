package domain

import (
	"github.com/google/uuid"
	"time"
)

type NotificationType int

const (
	NotificationTypeEmail NotificationType = iota
)

type Notification struct {
	Id        string
	Addressee string
	Message   string
	Type      NotificationType
	CreatedAt time.Time
}

func CreateNotification(addressee, msg string, t NotificationType) *Notification {
	return &Notification{
		Id:        uuid.NewString(),
		Addressee: addressee,
		Message:   msg,
		Type:      t,
		CreatedAt: time.Now(),
	}
}
