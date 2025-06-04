package api

import "github.com/canonical/gomaasclient/entity"

// Notifications is an interface for listing and creating Notification objects
type Notifications interface {
	Get() ([]entity.Notification, error)
	Create(params *entity.NotificationParams) (*entity.Notification, error)
}
