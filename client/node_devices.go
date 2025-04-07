package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// NodeDevices implements api.NodeDevices
type NodeDevices struct {
	APIClient APIClient
}

func (a *NodeDevices) client(systemID string) *APIClient {
	return a.APIClient.SubClient("nodes").
		SubClient(fmt.Sprintf("%v", systemID)).
		SubClient("devices")
}

// Get fetches a list of NodeDevice objects
func (a *NodeDevices) Get(ctx context.Context, systemID string, param *entity.NodeDeviceParams) ([]entity.NodeDevice, error) {
	qsp, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	nodeDevices := make([]entity.NodeDevice, 0)
	err = a.client(systemID).Get(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, &nodeDevices)
	})

	return nodeDevices, err
}
