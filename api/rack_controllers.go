package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// RackControllers represents the MAAS Rack Controllers endpoint
type RackControllers interface {
	DescribePowerTypes() ([]entity.PowerType, error)
	IsRegistered(macAddress string) (bool, error)
	GetPowerParameters(systemIDs []string) (map[string]interface{}, error)
	Get(*entity.RackControllersGetParams) ([]entity.RackController, error)
	SetZone(*entity.RackControllerSetZoneParams) error
}
