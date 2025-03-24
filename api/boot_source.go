package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BootSource is an interface defining API behaviour for
// boot sources
type BootSource interface {
	Get(ctx context.Context, id int) (*entity.BootSource, error)
	Update(ctx context.Context, id int, params *entity.BootSourceParams) (*entity.BootSource, error)
	Delete(ctx context.Context, id int) error
}
