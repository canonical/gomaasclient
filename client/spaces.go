//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// Spaces implements api.Spaces
type Spaces struct {
	APIClient APIClient
}

func (s *Spaces) client() *APIClient {
	return s.APIClient.SubClient("spaces")
}

// Get fetches a list of Space objects
func (s *Spaces) Get(ctx context.Context) ([]entity.Space, error) {
	spaces := make([]entity.Space, 0)
	err := s.client().Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &spaces)
	})

	return spaces, err
}

// Create creates a new Space
func (s *Spaces) Create(ctx context.Context, name string) (*entity.Space, error) {
	space := new(entity.Space)
	qsp := url.Values{}
	qsp.Set("name", name)
	err := s.client().Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, space)
	})

	return space, err
}
