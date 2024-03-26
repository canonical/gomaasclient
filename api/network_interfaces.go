package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// NetworkInterfaces represents the MAAS Server Interfaces endpoint
type NetworkInterfaces interface {
	Get(systemID string) ([]entity.NetworkInterface, error)
	CreateBond(systemID string, params *entity.NetworkInterfaceBondParams) (*entity.NetworkInterface, error)
	CreateBridge(systemID string, params *entity.NetworkInterfaceBridgeParams) (*entity.NetworkInterface, error)
	CreatePhysical(systemID string, params *entity.NetworkInterfacePhysicalParams) (*entity.NetworkInterface, error)
	CreateVLAN(systemID string, params *entity.NetworkInterfaceVLANParams) (*entity.NetworkInterface, error)
}
