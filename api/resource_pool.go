package api

import (
	"github.com/maas/gomaasclient/entity"
)

type ResourcePool interface {
	Get(id int) (*entity.ResourcePool, error)
	Update(id int, params *entity.ResourcePoolParams) (*entity.ResourcePool, error)
	Delete(id int) error
}
