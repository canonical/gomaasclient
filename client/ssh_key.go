package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// SSHKey implements api.SSHKey
type SSHKey struct {
	APIClient APIClient
}

func (s *SSHKey) client(id int) *APIClient {
	return s.APIClient.SubClient("account/prefs/sshkeys").SubClient(fmt.Sprintf("%v", id))
}

// Get fetches a given SSHKey
func (s *SSHKey) Get(ctx context.Context, id int) (*entity.SSHKey, error) {
	sshKey := new(entity.SSHKey)
	err := s.client(id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, sshKey)
	})

	return sshKey, err
}

// Delete deletes a given SSHKey
func (s *SSHKey) Delete(ctx context.Context, id int) error {
	return s.client(id).Delete(ctx)
}
