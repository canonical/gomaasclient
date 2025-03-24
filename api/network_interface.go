package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// NetworkInterface represents the MAAS Server Interface endpoint
type NetworkInterface interface {
	Get(ctx context.Context, systemID string, id int) (*entity.NetworkInterface, error)
	Update(ctx context.Context, systemID string, id int, params *entity.NetworkInterfaceUpdateParams) (*entity.NetworkInterface, error)
	Delete(ctx context.Context, systemID string, id int) error
	Disconnect(ctx context.Context, systemID string, id int) (*entity.NetworkInterface, error)
	AddTag(ctx context.Context, systemID string, id int, tag string) (*entity.NetworkInterface, error)
	RemoveTag(ctx context.Context, systemID string, id int, tag string) (*entity.NetworkInterface, error)
	LinkSubnet(ctx context.Context, systemID string, id int, params *entity.NetworkInterfaceLinkParams) (*entity.NetworkInterface, error)
	UnlinkSubnet(ctx context.Context, systemID string, id int, linkID int) (*entity.NetworkInterface, error)
	SetDefaultGateway(ctx context.Context, systemID string, id int, linkID int) (*entity.NetworkInterface, error)
}
