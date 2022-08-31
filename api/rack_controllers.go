package api

import (
	"github.com/maas/gomaasclient/entity"
)

// RackControllers represents the MaaS Rack Controllers endpoint
type RackControllers interface {
	Get(*entity.RackControllerSearch) ([]entity.RackController, error)
}
