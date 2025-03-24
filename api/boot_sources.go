package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BootSources is an interface for listing and creating
// BootSource objects
type BootSources interface {
	Get(ctx context.Context) ([]entity.BootSource, error)
	Create(ctx context.Context, params *entity.BootSourceParams) (*entity.BootSource, error)
}
