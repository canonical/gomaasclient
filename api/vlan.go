package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// VLAN is an interface defining API behaviour for VLAN objects
type VLAN interface {
	Get(ctx context.Context, fabricID int, vid int) (*entity.VLAN, error)
	Update(ctx context.Context, fabricID int, vid int, params *entity.VLANParams) (*entity.VLAN, error)
	Delete(ctx context.Context, fabricID int, vid int) error
}
