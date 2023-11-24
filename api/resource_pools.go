package api

import "github.com/maas/gomaasclient/entity"

type ResourcePools interface {
	Get() ([]entity.ResourcePool, error)
	Create(params *entity.ResourcePoolParams) (*entity.Domain, error)
}
