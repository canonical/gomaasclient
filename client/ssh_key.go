package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// SSHKey implements api.SSHKey
type SSHKey struct {
	APIClient APIClient
}

func (s *SSHKey) client(id int) APIClient {
	return s.APIClient.GetSubObject("account/prefs/sshkeys").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a given SSHKey
func (s *SSHKey) Get(id int) (*entity.SSHKey, error) {
	sshKey := new(entity.SSHKey)
	err := s.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, sshKey)
	})

	return sshKey, err
}

// Delete deletes a given SSHKey
func (s *SSHKey) Delete(id int) error {
	return s.client(id).Delete()
}
