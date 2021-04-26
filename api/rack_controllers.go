package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
)

// RackControllers represents the MaaS Rack Controllers endpoint
type RackControllers interface {
	Get(*endpoint.RackControllerSearch) ([]endpoint.RackController, error)
}
