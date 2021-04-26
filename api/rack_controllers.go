package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/params"
	"github.com/ionutbalutoiu/gomaasclient/maas/entity"
)

// RackControllers represents the MaaS Rack Controllers endpoint
type RackControllers interface {
	Get(*params.RackControllerSearch) ([]entity.RackController, error)
}
