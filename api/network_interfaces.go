package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/params"
	"github.com/ionutbalutoiu/gomaasclient/maas/entity"
)

// NetworkInterfaces represents the MaaS Server Interfaces endpoint
type NetworkInterfaces interface {
	Get(systemID string) ([]entity.NetworkInterface, error)
	CreateBond(systemID string, params *params.NetworkInterfaceBond) (*entity.NetworkInterface, error)
	CreateBridge(systemID string, params *params.NetworkInterfaceBridge) (*entity.NetworkInterface, error)
	CreatePhysical(systemID string, params *params.NetworkInterfacePhysical) (*entity.NetworkInterface, error)
	CreateVLAN(systemID string, params *params.NetworkInterfaceVLAN) (*entity.NetworkInterface, error)
}
