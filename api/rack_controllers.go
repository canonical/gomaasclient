package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// RackControllers represents the MAAS Rack Controllers endpoint
type RackControllers interface {
	Get(*entity.RackControllerSearch) ([]entity.RackController, error)
}
