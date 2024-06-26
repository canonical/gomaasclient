package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// ResourcePool is an interface defining API behaviour for ResourcePool
// objects
type ResourcePool interface {
	Get(id int) (*entity.ResourcePool, error)
	Update(id int, params *entity.ResourcePoolParams) (*entity.ResourcePool, error)
	Delete(id int) error
	GetByName(name string) (*entity.ResourcePool, error)
	UpdateByName(name string, params *entity.ResourcePoolParams) (*entity.ResourcePool, error)
	DeleteByName(name string) error
}
