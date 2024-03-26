package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// NodeDevice implements api.NodeDevice
type NodeDevice struct {
	APIClient APIClient
}

func (a *NodeDevice) client(systemID string, id int) APIClient {
	return a.APIClient.GetSubObject("nodes").
		GetSubObject(fmt.Sprintf("%v", systemID)).
		GetSubObject("devices").
		GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches NodeDevice object with id for given systemID
func (a *NodeDevice) Get(systemID string, id int) (*entity.NodeDevice, error) {
	nodeDevice := new(entity.NodeDevice)
	err := a.client(systemID, id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, nodeDevice)
	})

	return nodeDevice, err
}

// Delete deletes NodeDevice object with id for given systemID
func (a *NodeDevice) Delete(systemID string, id int) error {
	return a.client(systemID, id).Delete()
}
