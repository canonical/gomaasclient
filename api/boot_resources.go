package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BootResources is an interface for listing and creating
// BootResource objects
type BootResources interface {
	Get(ctx context.Context, params *entity.BootResourcesReadParams) ([]entity.BootResource, error)
	Create(ctx context.Context, params *entity.BootResourceParams) (*entity.BootResource, error)
	Import(ctx context.Context) error
	IsImporting(ctx context.Context) (bool, error)
	StopImport(ctx context.Context) error
}
