package client

import (
	"encoding/json"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Events implements api.Events
type Events struct {
	APIClient APIClient
}

func (e *Events) client() APIClient {
	return e.APIClient.GetSubObject("events")
}

// Get events for nodes
func (e *Events) Get(params *entity.EventParams) (*entity.EventsResp, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	events := &entity.EventsResp{}
	err = e.client().Get("query", qsp, func(data []byte) error {
		return json.Unmarshal(data, events)
	})

	return events, err
}
