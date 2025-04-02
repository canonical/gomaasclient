package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Fabric is an interface defining API behaviour for
// fabric objects
type Fabric interface {
	Get(ctx context.Context, id int) (*entity.Fabric, error)
	Update(ctx context.Context, id int, fabricParams *entity.FabricParams) (*entity.Fabric, error)
	Delete(ctx context.Context, id int) error
}
