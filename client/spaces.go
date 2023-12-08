package client

import (
	"encoding/json"
	"net/url"

	"github.com/maas/gomaasclient/entity"
)

// Spaces implements api.Spaces
type Spaces struct {
	ApiClient ApiClient
}

func (s *Spaces) client() ApiClient {
	return s.ApiClient.GetSubObject("spaces")
}

// Get fetches a list of Space objects
func (s *Spaces) Get() (spaces []entity.Space, err error) {
	err = s.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &spaces)
	})

	return
}

// Create creates a new Space
func (s *Spaces) Create(name string) (space *entity.Space, err error) {
	space = new(entity.Space)
	qsp := url.Values{}
	qsp.Set("name", name)
	err = s.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, space)
	})

	return
}
