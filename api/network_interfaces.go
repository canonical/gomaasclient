package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
)

// NetworkInterfaces represents the MaaS Server Interfaces endpoint
type NetworkInterfaces interface {
	Get(systemID string) ([]endpoint.NetworkInterface, error)
	CreateBond(systemID string, params *endpoint.NetworkInterfaceBondParams) (*endpoint.NetworkInterface, error)
	CreateBridge(systemID string, params *endpoint.NetworkInterfaceBridgeParams) (*endpoint.NetworkInterface, error)
	CreatePhysical(systemID string, params *endpoint.NetworkInterfacePhysicalParams) (*endpoint.NetworkInterface, error)
	CreateVLAN(systemID string, params *endpoint.NetworkInterfaceVLANParams) (*endpoint.NetworkInterface, error)
}
