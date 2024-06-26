package client

import (
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

func (n *NetworkInterface) client(systemID string, id int) APIClient {
	return n.APIClient.GetSubObject("nodes").
		GetSubObject(systemID).
		GetSubObject("interfaces").
		GetSubObject(fmt.Sprintf("%v", id))
}

func (n *NetworkInterface) callPost(systemID string, id int, qsp url.Values, op string) (*entity.NetworkInterface, error) {
	networkInterface := new(entity.NetworkInterface)
	err := n.client(systemID, id).Post(op, qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}

// Get fetches a NetworkInterface for the given system_id and NetworkInterface id
func (n *NetworkInterface) Get(systemID string, id int) (*entity.NetworkInterface, error) {
	networkInterface := new(entity.NetworkInterface)
	err := n.client(systemID, id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}

// Update updates a give NetworkInterface
func (n *NetworkInterface) Update(systemID string, id int, params *entity.NetworkInterfaceUpdateParams) (*entity.NetworkInterface, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	networkInterface := new(entity.NetworkInterface)
	err = n.client(systemID, id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})

	return networkInterface, err
}

// Delete deletes a given NetworkInterface
func (n *NetworkInterface) Delete(systemID string, id int) error {
	return n.client(systemID, id).Delete()
}

// Disconnect marks a given NetworkInterface as disconnected
func (n *NetworkInterface) Disconnect(systemID string, id int) (*entity.NetworkInterface, error) {
	return n.callPost(systemID, id, url.Values{}, "disconnect")
}

// AddTag adds a given tag to a given NetworkInterface
func (n *NetworkInterface) AddTag(systemID string, id int, tag string) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("tag", tag)

	return n.callPost(systemID, id, qsp, "add_tag")
}

// RemoveTag removes a given tag from a given NetworkInterface
func (n *NetworkInterface) RemoveTag(systemID string, id int, tag string) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("tag", tag)

	return n.callPost(systemID, id, qsp, "remove_tag")
}

// LinkSubnet links a given Subnet to a given NetworkInterface
func (n *NetworkInterface) LinkSubnet(systemID string, id int, params *entity.NetworkInterfaceLinkParams) (*entity.NetworkInterface, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	return n.callPost(systemID, id, qsp, "link_subnet")
}

// UnlinkSubnet removes a given link
func (n *NetworkInterface) UnlinkSubnet(systemID string, id int, linkID int) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("id", fmt.Sprintf("%v", linkID))

	return n.callPost(systemID, id, qsp, "unlink_subnet")
}

// SetDefaultGateway sets the default gateway of the given NetworkInterface
func (n *NetworkInterface) SetDefaultGateway(systemID string, id int, linkID int) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("link_id", fmt.Sprintf("%v", linkID))

	return n.callPost(systemID, id, qsp, "set_default_gateway")
}
