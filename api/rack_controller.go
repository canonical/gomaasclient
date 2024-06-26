package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// RackController represents the MAAS Rack Controller endpoint
type RackController interface {
	Get(systemID string) (*entity.RackController, error)
	Update(systemID string, machineParams *entity.RackControllerParams, powerParams map[string]interface{}) (*entity.RackController, error)
	Delete(systemID string) error
	GetPowerParameters(systemID string) (map[string]interface{}, error)
	PowerOn(systemID string, params *entity.RackControllerPowerOnParams) (*entity.RackController, error)
	PowerOff(systemID string, params *entity.RackControllerPowerOffParams) (*entity.RackController, error)
	GetPowerState(systemID string) (*entity.RackControllerPowerState, error)
	Abort(systemID string, comment string) (*entity.RackController, error)
	Details(systemID string) (*entity.RackControllerDetails, error)
	OverrideFailedTesting(systemID string, comment string) (*entity.RackController, error)
	Test(systemID string, params map[string]interface{}) (*entity.RackController, error)
}
