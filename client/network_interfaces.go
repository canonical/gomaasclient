package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type NetworkInterfaces struct {
	ApiClient ApiClient
}

func (n *NetworkInterfaces) client(systemID string) ApiClient {
	return n.ApiClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("interfaces")
}

func (n *NetworkInterfaces) Get(systemID string) (networkInterfaces []entity.NetworkInterface, err error) {
	err = n.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &networkInterfaces)
	})
	return
}

func (n *NetworkInterfaces) CreateBond(systemID string, params *entity.NetworkInterfaceBondParams) (networkInterface *entity.NetworkInterface, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID).Post("create_bond", qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}

func (n *NetworkInterfaces) CreateBridge(systemID string, params *entity.NetworkInterfaceBridgeParams) (networkInterface *entity.NetworkInterface, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID).Post("create_bridge", qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}

func (n *NetworkInterfaces) CreatePhysical(systemID string, params *entity.NetworkInterfacePhysicalParams) (networkInterface *entity.NetworkInterface, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID).Post("create_physical", qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}

func (n *NetworkInterfaces) CreateVLAN(systemID string, params *entity.NetworkInterfaceVLANParams) (networkInterface *entity.NetworkInterface, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID).Post("create_vlan", qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}
