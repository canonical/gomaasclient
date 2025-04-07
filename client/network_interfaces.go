package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// NetworkInterfaces implements api.NetworkInterfaces
type NetworkInterfaces struct {
	APIClient APIClient
}

func (n *NetworkInterfaces) client(systemID string) *APIClient {
	return n.APIClient.SubClient("nodes").SubClient(systemID).SubClient("interfaces")
}

// Get fetches a list of NetworkInterface objects
func (n *NetworkInterfaces) Get(ctx context.Context, systemID string) ([]entity.NetworkInterface, error) {
	networkInterfaces := make([]entity.NetworkInterface, 0)
	err := n.client(systemID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &networkInterfaces)
	})

	return networkInterfaces, err
}

// CreateBond creates a Bond interface from two or more given NetworkInterface objects
func (n *NetworkInterfaces) CreateBond(ctx context.Context, systemID string, params *entity.NetworkInterfaceBondParams) (*entity.NetworkInterface, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	networkInterface := new(entity.NetworkInterface)
	err = n.client(systemID).Post(ctx, "create_bond", qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}

// CreateBridge creates a Bridge type interface
func (n *NetworkInterfaces) CreateBridge(ctx context.Context, systemID string, params *entity.NetworkInterfaceBridgeParams) (*entity.NetworkInterface, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	networkInterface := new(entity.NetworkInterface)
	err = n.client(systemID).Post(ctx, "create_bridge", qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}

// CreatePhysical creates a new physical interface
func (n *NetworkInterfaces) CreatePhysical(ctx context.Context, systemID string, params *entity.NetworkInterfacePhysicalParams) (*entity.NetworkInterface, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	networkInterface := new(entity.NetworkInterface)
	err = n.client(systemID).Post(ctx, "create_physical", qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}

// CreateVLAN creates an interface tagged for a given VLAN
func (n *NetworkInterfaces) CreateVLAN(ctx context.Context, systemID string, params *entity.NetworkInterfaceVLANParams) (*entity.NetworkInterface, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	networkInterface := new(entity.NetworkInterface)
	err = n.client(systemID).Post(ctx, "create_vlan", qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}
