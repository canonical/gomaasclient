package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// VLANs represents the MAAS Vlans endpoint
type VLANs interface {
	Get(ctx context.Context, fabricID int) ([]entity.VLAN, error)
	Create(ctx context.Context, fabricID int, params *entity.VLANParams) (*entity.VLAN, error)
}
