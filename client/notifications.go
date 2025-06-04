//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Notifications implements api.Notifications
type Notifications struct {
	APIClient APIClient
}

func (n *Notifications) client() APIClient {
	return n.APIClient.GetSubObject("notifications")
}

// Get fetches a list of Notification objects
func (n *Notifications) Get() ([]entity.Notification, error) {
	notifications := make([]entity.Notification, 0)
	err := n.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &notifications)
	})

	return notifications, err
}

// Create creates a new Notification object
func (n *Notifications) Create(params *entity.NotificationParams) (*entity.Notification, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	notification := new(entity.Notification)
	err = n.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, notification)
	})

	return notification, err
}
