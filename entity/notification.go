package entity

import "github.com/canonical/gomaasclient/entity/notification"

type Notification struct {
	Context     map[string]interface{} `json:"context"`
	User        *User                  `json:"user"`
	ResourceURI string                 `json:"resource_uri"`
	Ident       string                 `json:"ident"`
	Message     string                 `json:"message"`
	Category    notification.Category  `json:"category"`
	ID          int                    `json:"id"`
	Users       bool                   `json:"users"`
	Admins      bool                   `json:"admins"`
	Dismissable bool                   `json:"dismissable"`
}

type NotificationParams struct {
	Context     map[string]interface{} `url:"context,omitempty"`
	Ident       string                 `url:"ident,omitempty"`
	Message     string                 `url:"message"`
	Category    notification.Category  `url:"category,omitempty"`
	User        int                    `url:"user,omitempty"`
	Users       bool                   `url:"users,omitempty"`
	Admins      bool                   `url:"admins,omitempty"`
	Dismissable bool                   `url:"dismissable,omitempty"`
}
