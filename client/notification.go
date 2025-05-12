package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Notification implements api.Notification
type Notification struct {
	APIClient APIClient
}

func (n *Notification) client(id int) APIClient {
	return n.APIClient.
		GetSubObject("notifications").
		GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a Notification object
func (n *Notification) Get(id int) (*entity.Notification, error) {
	notification := new(entity.Notification)
	err := n.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, notification)
	})

	return notification, err
}

// Update updates a given Notification
func (n *Notification) Update(id int, params *entity.NotificationParams) (*entity.Notification, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	notification := new(entity.Notification)
	err = n.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, notification)
	})

	return notification, err
}

// Delete deletes a given Notification
func (n *Notification) Delete(id int) error {
	return n.client(id).Delete()
}

// Dismiss dismisses a given Notification
func (n *Notification) Dismiss(id int) error {
	return n.client(id).Post("dismiss", url.Values{}, func(data []byte) error {
		return nil
	})
}
