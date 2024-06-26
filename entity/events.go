package entity

import "github.com/canonical/gomaasclient/entity/event"

// Event represents an event for MAAS node
type Event struct {
	UserName    string         `json:"username,omitempty"`
	Node        string         `json:"node,omitempty"`
	Hostname    string         `json:"hostname,omitempty"`
	Created     string         `json:"created,omitempty"`
	Type        string         `json:"type,omitempty"`
	Description string         `json:"description,omitempty"`
	Level       event.LogLevel `json:"level,omitempty"`
	ID          int            `json:"id,omitempty"`
}

// EventsResp represents the MAAS Events endpoint
type EventsResp struct {
	NextURI string  `json:"next_uri,omitempty"`
	PrevURI string  `json:"prev_uri,omitempty"`
	Events  []Event `json:"events,omitempty"`
	Count   int     `json:"count,omitempty"`
}

// EventParams enumerates the parameters for the event get operation
type EventParams struct {
	Hostname   string         `url:"hostname,omitempty"`
	MACAddress string         `url:"mac_address,omitempty"`
	ID         string         `url:"id,omitempty"`
	Zone       string         `url:"zone,omitempty"`
	AgentName  string         `url:"agent_name,omitempty"`
	Limit      string         `url:"limit,omitempty"`
	Before     string         `url:"before,omitempty"`
	After      string         `url:"after,omitempty"`
	Owner      string         `url:"owner,omitempty"`
	Level      event.LogLevel `url:"level,omitempty"`
}
