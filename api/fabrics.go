package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Fabrics is an interface for listing and creating Fabric records
type Fabrics interface {
	Get(ctx context.Context) ([]entity.Fabric, error)
	Create(ctx context.Context, fabricParams *entity.FabricParams) (*entity.Fabric, error)
}
