//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// SSLKeys implements api.SSLKeys
type SSLKeys struct {
	APIClient APIClient
}

func (s *SSLKeys) client() APIClient {
	return s.APIClient.GetSubObject("account/prefs/sslkeys")
}

// Get fetches a list of SSLKey objects
func (s *SSLKeys) Get() ([]entity.SSLKey, error) {
	sslKeys := make([]entity.SSLKey, 0)
	err := s.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &sslKeys)
	})

	return sslKeys, err
}

// Create creates a new SSLKey
func (s *SSLKeys) Create(key string) (*entity.SSLKey, error) {
	sslKey := new(entity.SSLKey)
	qsp := url.Values{}
	qsp.Set("key", key)
	err := s.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, sslKey)
	})

	return sslKey, err
}
