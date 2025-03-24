package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// ResourcePools is an interface for listing and creating ResourcePool records
type ResourcePools interface {
	Get(ctx context.Context) ([]entity.ResourcePool, error)
	Create(ctx context.Context, params *entity.ResourcePoolParams) (*entity.ResourcePool, error)
}
