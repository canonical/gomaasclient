package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// RackControllers represents the MAAS Rack Controllers endpoint
type RackControllers interface {
	DescribePowerTypes(ctx context.Context) ([]entity.PowerType, error)
	IsRegistered(ctx context.Context, macAddress string) (bool, error)
	GetPowerParameters(ctx context.Context, systemIDs []string) (map[string]interface{}, error)
	Get(ctx context.Context, params *entity.RackControllersGetParams) ([]entity.RackController, error)
	SetZone(ctx context.Context, params *entity.RackControllerSetZoneParams) error
}
