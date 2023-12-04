package api

import (
	"github.com/maas/gomaasclient/entity"
)

// ResourcePool is an interface defining API behaviour for ResourcePool
// objects
type ResourcePool interface {
	Get(id int) (*entity.ResourcePool, error)
	Update(id int, params *entity.ResourcePoolParams) (*entity.ResourcePool, error)
	Delete(id int) error
}
