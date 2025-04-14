//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// StaticRoutes implements api.StaticRoutes
type StaticRoutes struct {
	APIClient APIClient
}

func (s *StaticRoutes) client() APIClient {
	return s.APIClient.GetSubObject("static-routes")
}

// Get fetches a list of StaticRoutes objects
func (s *StaticRoutes) Get() ([]entity.StaticRoute, error) {
	staticRoutes := make([]entity.StaticRoute, 0)
	err := s.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &staticRoutes)
	})

	return staticRoutes, err
}

// Create creates a new StaticRoute
func (s *StaticRoutes) Create(params *entity.StaticRouteParams) (*entity.StaticRoute, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	staticRoute := new(entity.StaticRoute)
	err = s.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, staticRoute)
	})

	return staticRoute, err
}
