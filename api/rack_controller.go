package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// RackController represents the MAAS Rack Controller endpoint
type RackController interface {
	Get(ctx context.Context, systemID string) (*entity.RackController, error)
	Update(ctx context.Context, systemID string, machineParams *entity.RackControllerParams, powerParams map[string]interface{}) (*entity.RackController, error)
	Delete(ctx context.Context, systemID string) error
	GetPowerParameters(ctx context.Context, systemID string) (map[string]interface{}, error)
	PowerOn(ctx context.Context, systemID string, params *entity.RackControllerPowerOnParams) (*entity.RackController, error)
	PowerOff(ctx context.Context, systemID string, params *entity.RackControllerPowerOffParams) (*entity.RackController, error)
	GetPowerState(ctx context.Context, systemID string) (*entity.RackControllerPowerState, error)
	Abort(ctx context.Context, systemID string, comment string) (*entity.RackController, error)
	Details(ctx context.Context, systemID string) (*entity.RackControllerDetails, error)
	OverrideFailedTesting(ctx context.Context, systemID string, comment string) (*entity.RackController, error)
	Test(ctx context.Context, systemID string, params map[string]interface{}) (*entity.RackController, error)
}
