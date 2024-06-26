package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// SSLKey implements api.SSLKey
type SSLKey struct {
	APIClient APIClient
}

func (s *SSLKey) client(id int) APIClient {
	return s.APIClient.GetSubObject("account/prefs/sslkeys").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a given SSLKey
func (s *SSLKey) Get(id int) (*entity.SSLKey, error) {
	sslKey := new(entity.SSLKey)
	err := s.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, sslKey)
	})

	return sslKey, err
}

// Delete deletes a given SSLKey
func (s *SSLKey) Delete(id int) error {
	return s.client(id).Delete()
}
