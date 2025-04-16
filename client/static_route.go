package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// StaticRoute implements api.StaticRoute
type StaticRoute struct {
	APIClient APIClient
}

func (s *StaticRoute) client(id int) APIClient {
	return s.APIClient.GetSubObject("static-routes").GetSubObject(fmt.Sprintf("%v", id))
}

// Delete deletes a given StaticRoute
func (s *StaticRoute) Delete(id int) error {
	return s.client(id).Delete()
}

// Get fetches a given StaticRoute
func (s *StaticRoute) Get(id int) (*entity.StaticRoute, error) {
	staticRoute := new(entity.StaticRoute)
	err := s.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &staticRoute)
	})

	return staticRoute, err
}

// Update updates the given StaticRoute
func (s *StaticRoute) Update(id int, params *entity.StaticRouteParams) (*entity.StaticRoute, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	staticRoute := new(entity.StaticRoute)
	err = s.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, staticRoute)
	})

	return staticRoute, err
}
