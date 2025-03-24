package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// ResourcePool is an interface defining API behaviour for ResourcePool
// objects
type ResourcePool interface {
	Get(ctx context.Context, id int) (*entity.ResourcePool, error)
	Update(ctx context.Context, id int, params *entity.ResourcePoolParams) (*entity.ResourcePool, error)
	Delete(ctx context.Context, id int) error
	GetByName(ctx context.Context, name string) (*entity.ResourcePool, error)
	UpdateByName(ctx context.Context, name string, params *entity.ResourcePoolParams) (*entity.ResourcePool, error)
	DeleteByName(ctx context.Context, name string) error
}
