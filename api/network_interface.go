package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
)

// NetworkInterface represents the MaaS Server Interface endpoint
type NetworkInterface interface {
	Delete(systemID string, id int) error
	Get(systemID string, id int) (*endpoint.NetworkInterface, error)
	AddTag(systemID string, id int, tag string) (*endpoint.NetworkInterface, error)
	Disconnect(systemID string, id int) (*endpoint.NetworkInterface, error)
	LinkSubnet(systemID string, id int, params *endpoint.NetworkInterfaceLinkParams) (*endpoint.NetworkInterface, error)
	RemoveTag(systemID string, id int, tag string) (*endpoint.NetworkInterface, error)
	SetDefaultGateway(systemID string, id, linkID int) (*endpoint.NetworkInterface, error)
	UnlinkSubnet(systemID string, id, linkID int) (*endpoint.NetworkInterface, error)
	Put(systemID string, id int, params interface{}) (*endpoint.NetworkInterface, error)
}
