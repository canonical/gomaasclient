package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// SSHKeys implements api.SSHKeys
type SSHKeys struct {
	APIClient APIClient
}

func (s *SSHKeys) client() APIClient {
	return s.APIClient.GetSubObject("account/prefs/sshkeys")
}

// Get fetches a list of SSHKey objects
func (s *SSHKeys) Get() ([]entity.SSHKey, error) {
	sshKeys := make([]entity.SSHKey, 0)
	err := s.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &sshKeys)
	})

	return sshKeys, err
}

// Create creates a new SSHKey
func (s *SSHKeys) Create(key string) (*entity.SSHKey, error) {
	sshKey := new(entity.SSHKey)
	qsp := url.Values{}
	qsp.Set("key", key)
	err := s.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, sshKey)
	})

	return sshKey, err
}

// Import creates a new SSHKey
func (s *SSHKeys) Import(keysource string) ([]entity.SSHKey, error) {
	sshKeys := make([]entity.SSHKey, 0)
	qsp := url.Values{}
	qsp.Set("keysource", keysource)
	err := s.client().Post("import", qsp, func(data []byte) error {
		return json.Unmarshal(data, &sshKeys)
	})

	return sshKeys, err
}
