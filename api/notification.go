package api

import "github.com/canonical/gomaasclient/entity"

// Notification is an interface defining API behavior for
// Notification objects
type Notification interface {
	Get(id int) (*entity.Notification, error)
	Update(id int, params *entity.NotificationParams) (*entity.Notification, error)
	Delete(id int) error
	Dismiss(id int) error
}
