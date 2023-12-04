package api

import "github.com/maas/gomaasclient/entity"

// ResourcePools is an interface for listing and creating ResourcePool records
type ResourcePools interface {
	Get() ([]entity.ResourcePool, error)
	Create(params *entity.ResourcePoolParams) (*entity.Domain, error)
}
