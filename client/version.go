package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// Version implements api.Version
type Version struct {
	APIClient APIClient
}

func (v *Version) client() APIClient {
	return v.APIClient.GetSubObject("version")
}

// Get fetches MAAS version details
func (v *Version) Get() (*entity.Version, error) {
	version := new(entity.Version)
	err := v.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, version)
	})

	return version, err
}
