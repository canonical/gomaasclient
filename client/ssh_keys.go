package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// SSHKeys implements api.SSHKeys
type SSHKeys struct {
	APIClient APIClient
}

func (s *SSHKeys) client() *APIClient {
	return s.APIClient.SubClient("account/prefs/sshkeys")
}

// Get fetches a list of SSHKey objects
func (s *SSHKeys) Get(ctx context.Context) ([]entity.SSHKey, error) {
	sshKeys := make([]entity.SSHKey, 0)
	err := s.client().Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &sshKeys)
	})

	return sshKeys, err
}

// Create creates a new SSHKey
func (s *SSHKeys) Create(ctx context.Context, key string) (*entity.SSHKey, error) {
	sshKey := new(entity.SSHKey)
	qsp := url.Values{}
	qsp.Set("key", key)
	err := s.client().Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, sshKey)
	})

	return sshKey, err
}

// Import creates a new SSHKey
func (s *SSHKeys) Import(ctx context.Context, keysource string) ([]entity.SSHKey, error) {
	sshKeys := make([]entity.SSHKey, 0)
	qsp := url.Values{}
	qsp.Set("keysource", keysource)
	err := s.client().Post(ctx, "import", qsp, func(data []byte) error {
		return json.Unmarshal(data, &sshKeys)
	})

	return sshKeys, err
}
