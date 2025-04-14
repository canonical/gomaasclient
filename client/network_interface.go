package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// NetworkInterface implements api.NetworkInterface
type NetworkInterface struct {
	APIClient APIClient
}

func (n *NetworkInterface) client(systemID string, id int) *APIClient {
	return n.APIClient.SubClient("nodes").
		SubClient(systemID).
		SubClient("interfaces").
		SubClient(fmt.Sprintf("%v", id))
}

func (n *NetworkInterface) callPost(ctx context.Context, systemID string, id int, qsp url.Values, op string) (*entity.NetworkInterface, error) {
	networkInterface := new(entity.NetworkInterface)
	err := n.client(systemID, id).Post(ctx, op, qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}

// Get fetches a NetworkInterface for the given system_id and NetworkInterface id
func (n *NetworkInterface) Get(ctx context.Context, systemID string, id int) (*entity.NetworkInterface, error) {
	networkInterface := new(entity.NetworkInterface)
	err := n.client(systemID, id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}

// Update updates a give NetworkInterface
func (n *NetworkInterface) Update(ctx context.Context, systemID string, id int, params *entity.NetworkInterfaceUpdateParams) (*entity.NetworkInterface, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	networkInterface := new(entity.NetworkInterface)
	err = n.client(systemID, id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}

// Delete deletes a given NetworkInterface
func (n *NetworkInterface) Delete(ctx context.Context, systemID string, id int) error {
	return n.client(systemID, id).Delete(ctx)
}

// Disconnect marks a given NetworkInterface as disconnected
func (n *NetworkInterface) Disconnect(ctx context.Context, systemID string, id int) (*entity.NetworkInterface, error) {
	return n.callPost(ctx, systemID, id, url.Values{}, "disconnect")
}

// AddTag adds a given tag to a given NetworkInterface
func (n *NetworkInterface) AddTag(ctx context.Context, systemID string, id int, tag string) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("tag", tag)

	return n.callPost(ctx, systemID, id, qsp, "add_tag")
}

// RemoveTag removes a given tag from a given NetworkInterface
func (n *NetworkInterface) RemoveTag(ctx context.Context, systemID string, id int, tag string) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("tag", tag)

	return n.callPost(ctx, systemID, id, qsp, "remove_tag")
}

// LinkSubnet links a given Subnet to a given NetworkInterface
func (n *NetworkInterface) LinkSubnet(ctx context.Context, systemID string, id int, params *entity.NetworkInterfaceLinkParams) (*entity.NetworkInterface, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	return n.callPost(ctx, systemID, id, qsp, "link_subnet")
}

// UnlinkSubnet removes a given link
func (n *NetworkInterface) UnlinkSubnet(ctx context.Context, systemID string, id int, linkID int) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("id", fmt.Sprintf("%v", linkID))

	return n.callPost(ctx, systemID, id, qsp, "unlink_subnet")
}

// SetDefaultGateway sets the default gateway of the given NetworkInterface
func (n *NetworkInterface) SetDefaultGateway(ctx context.Context, systemID string, id int, linkID int) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("link_id", fmt.Sprintf("%v", linkID))

	return n.callPost(ctx, systemID, id, qsp, "set_default_gateway")
}
