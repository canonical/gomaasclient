package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// NetworkInterfaces represents the MAAS Server Interfaces endpoint
type NetworkInterfaces interface {
	Get(ctx context.Context, systemID string) ([]entity.NetworkInterface, error)
	CreateBond(ctx context.Context, systemID string, params *entity.NetworkInterfaceBondParams) (*entity.NetworkInterface, error)
	CreateBridge(ctx context.Context, systemID string, params *entity.NetworkInterfaceBridgeParams) (*entity.NetworkInterface, error)
	CreatePhysical(ctx context.Context, systemID string, params *entity.NetworkInterfacePhysicalParams) (*entity.NetworkInterface, error)
	CreateVLAN(ctx context.Context, systemID string, params *entity.NetworkInterfaceVLANParams) (*entity.NetworkInterface, error)
}
